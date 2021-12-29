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
	TemplateDir string = "server_side_app/views/"
	TemplateExt string = ".html"
)

type SsRoutes struct{}

type ssRoutes interface {
	NotFoundPage(w http.ResponseWriter, r *http.Request)
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

func NewView(layout string, files ...string) *View {
	var t *template.Template

	addTemplatePath(files)
	addTemplateExt((files))

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

func (route *SsRoutes) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Hi there! Not Found!</h1>")

}
