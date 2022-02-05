package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layout"

// NewView is responsible for handling all of the work required to prepare
// a view. Once the view is prepared, NewView returns a pointer to an object of
// type View.
func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	fmt.Println("call to NewView: ", files) // debugging
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
	fmt.Println("rendering: ", v.Layout, "with this data", data)
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles globs all of the files ending with extensions .gohtml
// in our views/layout directory. It returns a slice of strings where each
// string in the slice is the path to a layout file ending in .gohtml
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	fmt.Println(LayoutDir)
	fmt.Println("from layoutFiles", files)
	if err != nil {
		panic(err)
	}
	return files
}
