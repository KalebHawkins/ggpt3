package ggpt3

import (
	"context"
	"net/http"
)

func (c *Client) RequestEmbedding(ctx context.Context, er *EmbeddingRequest) (*EmbeddingsResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, nil, "/embeddings", er)
	if err != nil {
		return nil, err
	}

	var eResp EmbeddingsResponse
	err = c.sendRequest(req, &eResp)

	return &eResp, err
}
