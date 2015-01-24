package pages

import (
	"bitodd/util"
	"net/http"
)

const indexUrl = "/"

var indexTmpl = getTemplate("templates/index.html")

// Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmplModel := &util.TemplateModel{Params: make(map[interface{}]interface{}, 0)}

	util.RenderTemplate(indexTmpl, w, tmplModel)
}
