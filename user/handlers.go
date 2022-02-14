package user

import (
	"net/http"
	"quiltspace/config"
	"quiltspace/views"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// SignUpForm renders the template to display a form to submit information about
// a new user profile. Once the submit is clicked, a call to SignUpProcess actually
// does the database processing and user creation.
func SignUpForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	views.SignUpFormView.Render(w, nil)
}

// SignUpProcess does the database transactions for creating a user, encrypting
// their password, and rendering a template that tells them everything was
// created. It won't log them in to their account after completion.
func SignUpProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	usr := User{}
	usr.Email = r.FormValue("email")
	usr.Username = r.FormValue("username")
	usr.Password = r.FormValue("password")
	usr.First = r.FormValue("firstname")
	usr.Last = r.FormValue("lastname")

	// validate form values
	if usr.Email == "" || usr.Username == "" || usr.Password == "" || usr.First == "" || usr.Last == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values into database
	bs, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (email, username, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5)", usr.Email, usr.Username, bs, usr.First, usr.Last)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	views.SignUpFormProcessView.Render(w, usr)

}

func LoginForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	views.LoginFormView.Render(w, nil)
}
