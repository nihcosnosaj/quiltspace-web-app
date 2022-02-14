package main

import (
	"net/http"
	"quiltspace/quilts"
	"quiltspace/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// handling routing with third-party router from httprouter
	router := httprouter.New()
	router.GET("/", quilts.Home)
	router.GET("/quilts", quilts.Index)
	router.GET("/about", quilts.About)
	router.GET("/quilts/show", quilts.Show)
	router.GET("/quilts/create", quilts.CreateForm)
	router.POST("/quilts/create/process", quilts.CreateProcess)
	router.GET("/quilts/update", quilts.UpdateForm)
	router.POST("/quilts/update/process", quilts.UpdateProcess)
	router.GET("/quilts/delete/process", quilts.DeleteProcess)
	router.GET("/user/create", user.SignUpForm)
	router.POST("/user/create/process", user.SignUpProcess)
	router.GET("/user/login", user.LoginForm)
	http.ListenAndServe(":8080", router)
}
