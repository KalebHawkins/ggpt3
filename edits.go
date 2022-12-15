package ggpt3

import (
	"context"
	"net/http"
)

func (c *Client) RequestEdits(ctx context.Context, er *EditsRequest) (*EditsResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json; charset=utf-8")

	req, err := c.newRequest(ctx, http.MethodPost, headers, "/edits", er)
	if err != nil {
		return nil, err
	}

	var eResp EditsResponse
	err = c.sendRequest(req, &eResp)
	return &eResp, err
}
