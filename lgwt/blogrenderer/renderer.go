package blogrenderer

import (
	"fmt"
	"io"

	posts "msw/lgwt/blogposts"
)

func Render(w io.Writer, p posts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
