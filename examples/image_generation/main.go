//go:build example
// +build example

package main

import (
	"context"
	_ "embed"
	"encoding/base64"
	"fmt"
	_ "image/png"
	"os"

	"github.com/KalebHawkins/ggpt3"
)

func main() {
	c := ggpt3.NewClient(os.Getenv("API_KEY"))

	ir := &ggpt3.ImageRequest{
		Prompt:         "A chicken with its head cut off",
		N:              1,
		Size:           "512x512",
		ResponseFormat: "b64_json",
	}

	imgResp, err := c.RequestImages(context.Background(), ir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	encodedImg := imgResp.Data[0].B64
	imgBuf, err := base64.StdEncoding.DecodeString(encodedImg)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	err = os.WriteFile("img.png", imgBuf, os.ModeAppend)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
