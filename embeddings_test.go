package ggpt3

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedding(t *testing.T) {
	tcs := []struct {
		name    string
		apiKey  string
		expResp *EmbeddingsResponse
		expErr  error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			expResp: &EmbeddingsResponse{
				Object: "list",
				Data: []data{
					{
						Object:    "embedding",
						Embedding: []float64{0.01899, -0.00738, 0.02127},
						Index:     0,
					},
				},
				Usage: usage{
					PromptTokens: 8,
					TotalTokens:  8,
				},
			},
			expErr: nil,
		},
		{
			name:    "WithErr",
			apiKey:  "",
			expResp: nil,
			expErr:  errApiKeyUndefined,
		},
	}

	ts := startTestServer(t)
	defer ts.Close()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(testApiKey)
			c.baseUrl = ts.URL

			if tc.expErr != nil {
				c.apiKey = ""
			}

			eReq := &EmbeddingRequest{
				Model: TextSimilarityBabbage001,
				Input: []string{"The food was delicious and the waiter..."},
			}
			embedResp, err := c.RequestEmbedding(context.TODO(), eReq)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expResp, embedResp)
		})
	}
}
