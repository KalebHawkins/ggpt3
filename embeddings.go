package ggpt3

import (
	"context"
	"net/http"
)

func (c *Client) RequestEmbedding(ctx context.Context, er *EmbeddingRequest) (*EmbeddingsResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json; charset=utf-8")

	req, err := c.newRequest(ctx, http.MethodPost, headers, "/embeddings", er)
	if err != nil {
		return nil, err
	}

	var eResp EmbeddingsResponse
	err = c.sendRequest(req, &eResp)

	return &eResp, err
}
