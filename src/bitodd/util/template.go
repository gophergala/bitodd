package util

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateModel struct {
	Params   map[interface{}]interface{}
}

func noescape(text string) template.HTML {
	return template.HTML(text)
}

func GetTemplateBase(name string) *template.Template {

	f := template.FuncMap{"Noescape": noescape}
	t := template.New(filepath.Base(name)).Funcs(f)

	return t
}

func GetNamedTemplate(name string) *template.Template {
	return template.Must(GetTemplateBase(name).ParseFiles(name))
}

func RenderTemplate(tmpl *template.Template, w http.ResponseWriter, model *TemplateModel) {

	err := tmpl.Execute(w, model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Redirect
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusFound)
}
