package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the Quiltspace!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	http.ListenAndServe(":8080", router)
}
