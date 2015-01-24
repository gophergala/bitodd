package pages

import (
	"bitodd/util"
	"github.com/gorilla/mux"
	"html/template"
)

func getTemplate(filename string) *template.Template {
	return template.Must(util.GetNamedTemplate("templates/base.html").ParseFiles(filename))
}

func GetRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc(indexUrl, indexHandler)

	return r
}
