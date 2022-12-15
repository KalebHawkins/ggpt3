package ggpt3

import (
	"context"
	"net/http"
)

// RequestCompletion creates a completion for the provided prompt and parameters
func (c *Client) RequestCompletion(ctx context.Context, cr *CompletionRequest) (*CompletionResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json; charset=utf-8")

	req, err := c.newRequest(ctx, http.MethodPost, headers, "/completions", cr)
	if err != nil {
		return nil, err
	}

	var cResp CompletionResponse
	err = c.sendRequest(req, &cResp)

	return &cResp, err
}
