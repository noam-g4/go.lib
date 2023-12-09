package handler

import (
	"os"
	"strings"

	"go.uber.org/fx"
)

type HandlersGenerator struct {
	handlers []Handler
}

type deps struct {
	fx.In

	Templates Templates
}

func NewHandlersGenerator(d deps) (*HandlersGenerator, error) {
	endpoints, err := getHtmlFileNames("./content")
	if err != nil {
		return nil, err
	}

	handlers := make([]Handler, 0)
	for _, endpoint := range endpoints {
		h, err := newHandler(endpoint, d.Templates)
		if err != nil {
			return nil, err
		}

		handlers = append(handlers, h)
	}
	return &HandlersGenerator{
		handlers: handlers,
	}, nil
}

func (h *HandlersGenerator) Handlers() []Handler {
	return h.handlers
}

func getHtmlFileNames(path string) ([]string, error) {
	output := make([]string, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		return output, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		if strings.HasSuffix(strings.ToLower(file.Name()), ".html") {
			output = append(output, file.Name())
		}
	}

	return output, nil
}
