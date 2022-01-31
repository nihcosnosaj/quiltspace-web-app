package main

import (
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
	http.ListenAndServe(":8080", router)
}

func quiltsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// The quiltsIndex handler will handle all requests to "/" and "/quilts".
	// It will display a formatted list of all quilts in the database.
	rows, err := config.DB.Query("SELECT * FROM quilts;")
	if err != nil {
		log.Fatalln(err)
	}

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

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/quilts", http.StatusSeeOther)
}
