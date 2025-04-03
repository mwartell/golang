package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	posts "lgwt/blogposts"
	"lgwt/blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = posts.Post{
			Title:       "hello world",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
			Body:		"A post body with some **markdown** highlighting and [a link](https://example.com)",
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
func BenchmarkRender(b *testing.B) {
	var (
		aPost = posts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}


	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
