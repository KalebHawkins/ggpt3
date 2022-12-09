package ggpt3

import (
	"context"
	"net/http"
)

// RequestCompletion creates a completion for the provided prompt and parameters
func (c *Client) RequestCompletion(ctx context.Context, cr *CompletionRequest) (*CompletionResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, nil, "/completions", cr)
	if err != nil {
		return nil, err
	}

	var cResp CompletionResponse
	err = c.sendRequest(req, &cResp)

	return &cResp, err
}
