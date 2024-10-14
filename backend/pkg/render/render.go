package render

import "html/template"

func NewRender() (*template.Template, error) {
	return template.ParseFiles(
		"public/index.html",
		"public/websocket-msg.html",
		"public/disk.html",
		"public/snapraid.html",
		"public/system.html",
	)
}
