package handler

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Templates interface {
	Get() map[string]template.HTML
}

type templates struct {
	m map[string]template.HTML
}

func NewTemplates() (Templates, error) {
	templateFileNames, err := getHtmlFileNames("./content/templates")
	if err != nil {
		return nil, err
	}

	temps := make(map[string]template.HTML)
	for _, t := range templateFileNames {
		data, err := os.ReadFile(fmt.Sprintf("./content/templates/%s", t))
		if err != nil {
			return nil, err
		}

		key := fmt.Sprintf("template_%s", strings.TrimSuffix(strings.ToLower(t), ".html"))
		temps[key] = template.HTML(data)
	}

	return &templates{m: temps}, nil
}

func (t *templates) Get() map[string]template.HTML {
	return t.m
}
