package main

import (
	"database/sql"
	"log"
	"net/http"
	"quiltspace/config"

	"github.com/julienschmidt/httprouter"
)

type Quilt struct {
	Qid     int
	Name    string
	Style   string
	Pattern string
}

func main() {

	// handling routing with router from httprouter
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/quilts", quiltsIndex)
	router.GET("/quilts/show", quiltsShow)
	router.GET("/quilts/create", quiltsCreateForm)
	router.POST("/quilts/create/process", quiltsCreateProcess)
	http.ListenAndServe(":8080", router)
}

// quiltsIndex handles all requests to "/" and "/quilts" from the client. It displays an index
// of all quilts in the database. It querys for all quilts, then creates a slice
// of type Quilt in which to place each quilt after the row has been scanned and
// partitioned accordingly.
func quiltsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// query database for all quilts
	rows, err := config.DB.Query("SELECT * FROM quilts;")
	if err != nil {
		log.Fatalln(err)
	}

	// Scan each row of query and add each row to a slice of type Quilt
	qlts := make([]Quilt, 0)
	for rows.Next() {
		qlt := Quilt{}
		err := rows.Scan(&qlt.Qid, &qlt.Name, &qlt.Style, &qlt.Pattern) // order matters here!
		if err != nil {
			log.Fatal(err)
		}
		qlts = append(qlts, qlt)
	}

	config.TPL.ExecuteTemplate(w, "index.html", qlts)
}

// quiltsShow handles all requests to "/quilts/show" from the client. It displays a
// single quilt based on the passed in quilt name, "name", from the FormValue of the
// request.
func quiltsShow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// retrieve quilt name queried from request & make sure it isn't empty
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// query the database for the requested quilt
	row := config.DB.QueryRow("SELECT * FROM quilts WHERE name = $1", name)

	// marshal query into type Quilt data structure for passing to template
	qlt := Quilt{}
	err := row.Scan(&qlt.Qid, &qlt.Name, &qlt.Style, &qlt.Pattern)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.html", qlt)

}

// quiltsCreateForm simply brings up the form for entering information for a new
// quilting project to be entered into the database. Once the "submit" button is
// entered on the webpage, a call to quiltsCreateProcess actually processes the
// insertion into the database.
func quiltsCreateForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "create.html", nil)
}

// quiltsCreateProcess handles inserting the information the user inputted into the form
// on the webpage and validates the data. It then inserts it into the database and executes
// a confirmation template to let the user know the quilt was successfully created.
func quiltsCreateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	qlt := Quilt{}
	qlt.Name = r.FormValue("name")
	qlt.Style = r.FormValue("style")
	qlt.Pattern = r.FormValue("pattern")

	// validate form values
	if qlt.Name == "" || qlt.Style == "" || qlt.Pattern == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values into database
	var err error
	_, err = config.DB.Exec("INSERT INTO quilts (name, style, pattern) VALUES ($1, $2, $3)", qlt.Name, qlt.Style, qlt.Pattern)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion into database
	config.TPL.ExecuteTemplate(w, "created.html", qlt)
}

// index redirects all requests to "/" to "/quilts"
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/quilts", http.StatusSeeOther)
}
