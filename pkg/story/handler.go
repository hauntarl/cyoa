package story

import (
	"html/template"
	"log"
	"net/http"
)

// HandlerOption allows users to provide multiple options while
// creating a new http.Handler
type HandlerOption func(h *handler)

// NewHandler takes in a story and returns the http.Handler for that story
func NewHandler(s Book, opts ...HandlerOption) http.Handler {
	h := handler{s, tmpl, defaultPathFn} // creating a default handler
	for _, opt := range opts {
		opt(&h)
	}
	return &h
}

type handler struct {
	s      Book
	t      *template.Template
	pathFn func(r *http.Request) string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	chapter, ok := h.s[path]
	if !ok {
		http.Error(w, "Chapter not found", http.StatusNotFound)
		return
	}
	err := h.t.Execute(w, chapter)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}
