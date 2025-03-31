package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// By default, GenerateRequest is streaming.
	req := &api.GenerateRequest{
		Model: "gemma3:4b",
		System: `You are a markdown blog post writer.
		   Write a blog post about the topic below.
		   At the top of the post should be metadata in the following format:
		   ---
		   title: <title>
		   description: <description>
		   tags: <tag1>, <tag2>, <tag3>
		   ---
		   The post body should be in markdown format, and should contain no
		   more than 500 words. Do not ask questions at the end of the post.`,
		Prompt: "Write a post about the structure of US federal government.",
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.

		// In streaming mode, responses are partial so we call fmt.Print (and not
		// Println) in order to avoid spurious newlines being introduced. The
		// model will insert its own newlines if it wants.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
