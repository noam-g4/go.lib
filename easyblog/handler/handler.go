package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Handler interface {
	http.Handler
	Pattern() string
}

type handler struct {
	templatesEngine Templates
	template        *template.Template
	endpoint        string
}

func newHandler(filename string, templates Templates) (Handler, error) {
	endpoint := resolveEndpoint(filename)
	template, err := loadTemplate(filename)
	if err != nil {
		return nil, err
	}
	return &handler{
		endpoint:        endpoint,
		template:        template,
		templatesEngine: templates,
	}, nil
}

func (h *handler) Pattern() string {
	return h.endpoint
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.template.Execute(w, h.templatesEngine.Get()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func resolveEndpoint(name string) string {
	endpoint := fmt.Sprintf("/%s",
		strings.ToLower(strings.TrimSuffix(name, ".html")))
	if strings.ToLower(name) == "index.html" {
		endpoint = "/"
	}

	return endpoint
}

func loadTemplate(filename string) (*template.Template, error) {
	path := fmt.Sprintf("./content/%s", filename)
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return template.Must(template.New(filename).Parse(string(file))), nil
}
