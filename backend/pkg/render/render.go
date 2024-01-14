package render

import "html/template"

func NewRender() (*template.Template, error) {
	return template.ParseFiles("public/index.html")
}
