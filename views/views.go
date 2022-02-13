package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layout"

var IndexView *View
var ShowView *View
var CreateFormView *View
var CreateProcessView *View
var UpdateFormView *View
var UpdateProcessView *View
var HomeView *View
var AboutView *View
var SignUpFormView *View
var SignUpFormProcessView *View

func init() {
	// intialiaze new views a
	HomeView = NewView("bootstrap", "views/home.gohtml")
	IndexView = NewView("bootstrap", "views/index.gohtml")
	ShowView = NewView("bootstrap", "views/show.gohtml")
	CreateFormView = NewView("bootstrap", "views/create.gohtml")
	CreateProcessView = NewView("bootstrap", "views/created.gohtml")
	UpdateFormView = NewView("bootstrap", "views/update.gohtml")
	UpdateProcessView = NewView("bootstrap", "views/updated.gohtml")
	AboutView = NewView("bootstrap", "views/about.gohtml")
	SignUpFormView = NewView("bootstrap", "views/signup.gohtml")
	SignUpFormProcessView = NewView("bootstrap", "views/signedup.gohtml")
}

// NewView is responsible for handling all of the work required to prepare
// a view. Once the view is prepared, NewView returns a pointer to an object of
// type View.
func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Render handles executing the template using the provided data interface. The results are sent to the response writer.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles globs all of the files ending with extensions .gohtml
// in our views/layout directory. It returns a slice of strings where each
// string in the slice is the path to a layout file ending in .gohtml
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
