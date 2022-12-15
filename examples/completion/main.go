package main

import (
	"context"
	"fmt"
	"os"

	"github.com/KalebHawkins/ggpt3"
)

func main() {
	c := ggpt3.NewClient(os.Getenv("API_KEY"))

	cr := &ggpt3.CompletionRequest{
		Model:     ggpt3.TextDavinci003,
		Prompt:    "Say this is a test.",
		MaxTokens: 150,
	}

	resp, err := c.RequestCompletion(context.Background(), cr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Your image was generated at: %s\n", resp.Choices[0].Text)
}
