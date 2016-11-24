package msrvd

import (
	"github.com/msawangwan/mhttplib/mtngin"
	"log"
	"net/http"
)

const ( // todo: add cli capability
	DIR_TEMPLATES_ROOT = "../mtngin/template/"
	DIR_LAYOUTS        = "layout/"
	DIR_STATICS        = "static/"
)

type StaticContentHandler struct {
	*mtngin.TemplateEngine
}

var sch *StaticContentHandler

func (sch *StaticContentHandler) ServeStaticContent(port string) {
	http.ListenAndServe(port, nil)
}

func NewStaticContentHandler() *StaticContentHandler {
	sch = &StaticContentHandler{
		TemplateEngine: mtngin.NewTemplateEngine(DIR_TEMPLATES_ROOT, DIR_LAYOUTS, DIR_STATICS),
	}

	layout, _ := sch.ParseDirTemplates(sch.Layouts)
	statics, _ := sch.ParseDirTemplates(sch.Statics)

	basicLayout, ok := layout["base"] // from the file: base.html

	if !ok {
		log.Printf("could not find basic layout named 'base' in the cache\n")
	}

	sch.CacheTemplates(basicLayout, statics)

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/favicon.ico", faviconhandler)

	return sch
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("served")
	sch.TemplateCache["index"].ExecuteTemplate(w, "layout", "")
}

func faviconhandler(w http.ResponseWriter, r *http.Request) {
	/* empty handler handles the browser requesting favicon.ico */
}
