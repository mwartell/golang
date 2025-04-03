package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	posts "lgwt/blogposts"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p posts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
