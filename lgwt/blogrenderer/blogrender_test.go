package blogrenderer_test

import (
	"bytes"
	"testing"

	posts "lgwt/blogposts"
	"lgwt/blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = posts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
