package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	Routes      ssRoutes = &SsRoutes{}
	err         error
	LayoutDir   string = "server_side_app/views/layouts/"
	TemplateExt string = ".html"
)

type SsRoutes struct{}

type ssRoutes interface {
	Home(w http.ResponseWriter, r *http.Request)
	Contact(w http.ResponseWriter, r *http.Request)
	Faq(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	NotFoundPage(w http.ResponseWriter, r *http.Request)
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

func NewView(layout string, files ...string) *View {
	var t *template.Template
	files = append(files, layoutFiles()...)
	t, err = template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (route *SsRoutes) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route Logic
	homeView := NewView("main", "server_side_app/views/home.html")

	err = homeView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func (route *SsRoutes) Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route Logic
	contactView := NewView("main", "server_side_app/views/contact.html")

	err = contactView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func (route *SsRoutes) Faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route logic
	faqView := NewView("main", "server_side_app/views/faq.html")

	err = faqView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func (route *SsRoutes) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route logic
	signUpView := NewView("main", "server_side_app/views/signup.html")

	err = signUpView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func (route *SsRoutes) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Hi there! Not Found!</h1>")

}
