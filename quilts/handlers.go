package quilts

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"quiltspace/config"
	"quiltspace/views"

	"github.com/julienschmidt/httprouter"
)

// Index handles all requests to "/" and "/quilts" from the client. It displays an index
// of all quilts in the database. It querys for all quilts, then creates a slice
// of type Quilt in which to place each quilt after the row has been scanned and
// partitioned accordingly.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	views.IndexView.Render(w, qlts)
}

// About handles all requests to "/about" from the client. It displays a single page
// that details any administrative information about the project, as well as any contact
// information.
func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	views.AboutView.Render(w, nil)
}

// Show handles all requests to "/quilts/show" from the client. It displays a
// single quilt based on the passed in quilt name, "name", from the FormValue of the
// request.
func Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	views.ShowView.Render(w, qlt)

}

// CreateForm simply brings up the form for entering information for a new
// quilting project to be entered into the database. Once the "submit" button is
// entered on the webpage, a call to quiltsCreateProcess actually processes the
// insertion into the database.
func CreateForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	views.CreateFormView.Render(w, nil)
}

// CreateProcess handles inserting the information the user inputted into the form
// on the webpage and validates the data. It then inserts it into the database and executes
// a confirmation template to let the user know the quilt was successfully created.
func CreateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// confirm creation to user
	views.CreateProcessView.Render(w, qlt)
}

// UpdateForm grabs the name of the quilt needing updating from the template form
// and queries the database for that quilt. It then creates a new type quilt instance that
// holds the query data for the update form in the template executed at the end of the function.
func UpdateForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM quilts WHERE name = $1", name)

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

	views.UpdateFormView.Render(w, qlt)
}

// UpdateProcess executes the updates made to the quilt selected in quiltsUpdateForm.
// Yes, it does use POST to update an already existing resource, but current HTML only supports
// GET and POST in forms. A workaround for this so PUT can be used is on the to-do list.
// For now, quiltsUpdateProcess retrieves the form values from the update form, validates them,
// and then executes the update in the database based on the Quilt ID.
func UpdateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "POST" {
		fmt.Println(r.Method)
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	qlt := Quilt{}
	qlt.Qid = r.FormValue("qid")
	qlt.Name = r.FormValue("name")
	qlt.Style = r.FormValue("style")
	qlt.Pattern = r.FormValue("pattern")

	// validate form values
	if qlt.Name == "" || qlt.Style == "" || qlt.Pattern == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// update values in database
	var err error
	_, err = config.DB.Exec("UPDATE quilts SET name = $1, style = $2, pattern = $3 WHERE qid=$4", qlt.Name, qlt.Style, qlt.Pattern, qlt.Qid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	views.UpdateProcessView.Render(w, qlt)
}

// DeleteProcess takes a given quilt ID number and deletes it from the database.
// It doesn't render a template, and instead returns the client back to the main index of
// all quilts. As of now, there is not "Are you sure you want to delete this?" but it may
// be something worth implementing in the near future.
func DeleteProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	quiltid := r.FormValue("qid")
	if quiltid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete quilt
	_, err := config.DB.Exec("DELETE FROM quilts WHERE qid=$1", quiltid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/quilts", http.StatusSeeOther)

}

// home displays the home page template and is the landing page for all requests
// to the root "/"
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	views.HomeView.Render(w, nil)
}
