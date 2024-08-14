package cyoa

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// handler is the main HTTP handler for the CYOA story.
// It implements the http.Handler interface and is responsible for rendering story chapters.
type handler struct {
	story  Story
	tmpl   *template.Template
	pathFn func(r *http.Request) string
}

// HandlerOption represents a functional option for configuring a Handler.
type HandlerOption func(h *handler)

// NewHandler creates a new Handler instance with the provided story and configuration options.
// It returns an http.Handler that can be used to serve HTTP requests.
//
// Example usage:
//
//	h := NewHandler(story, WithTemplate(tpl), WithPathFunc(customPathFn))
//	http.Handle("/story/", h)
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := &handler{
		story:  s,
		tmpl:   MustLoadTemplates("templates"), // Use defualt templates directory
		pathFn: DefaultPathFn,                  // Default path function to extract chapter names
	}
	// Apply each option to configure the handler
	for _, opt := range opts {
		opt(h)
	}

	return h
}

// WithTemplate returns a handlerOption that sets a custom template for rendering story chapters.
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tmpl = t
	}
}

// WithPathFunc returns a handlerOption that sets a custom function to extract the chapter name from the URL path.
func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

// ServeHTTP handles HTTP requests and renders the appropriate story chapter using the template.
// If the requested chapter is found in the story, it renders the chapter. Otherwise, it returns a 404 error.

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r) // Extract the chapter name from the URL path

	if chapter, ok := h.story[path]; ok {
		// Render the template with the chapter data
		err := h.tmpl.ExecuteTemplate(w, "story.html", chapter)
		if err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	// If the chapter is not found, return a 404 error
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// DefaultPathFn is the default function to extract the chapter name from the URL path.
// It strips leading slashes and defaults to "/intro" if no specific path is provided.
func DefaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:] // Remove leading '/'
}
