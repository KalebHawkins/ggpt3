package ggpt3

import (
	"context"
	"net/http"
)

func (c *Client) RequestEdits(ctx context.Context, er *EditsRequest) (*EditsResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, nil, "/edits", er)
	if err != nil {
		return nil, err
	}

	var eResp EditsResponse
	err = c.sendRequest(req, &eResp)
	return &eResp, err
}
