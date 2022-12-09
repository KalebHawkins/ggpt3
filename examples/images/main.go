package main

import (
	"context"
	_ "embed"
	"fmt"
	_ "image/png"
	"os"

	"github.com/KalebHawkins/ggpt3"
)

func main() {
	c := ggpt3.NewClient(os.Getenv("API_KEY"))

	ir := &ggpt3.ImageRequest{
		Image: "otter.png",
		N:     1,
		Size:  "500x500",
	}

	resp, err := c.RequestImageVariation(context.Background(), ir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Your image was generated at: %s\n", resp.Data[0].Url)
}
