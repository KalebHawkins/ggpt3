package ggpt3

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestCompletion(t *testing.T) {
	tcs := []struct {
		name    string
		apiKey  string
		expResp *CompletionResponse
		expErr  error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			expResp: &CompletionResponse{
				Id:      "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7",
				Object:  "text_completion",
				Created: 1589478378,
				Model:   "text-davinci-003",
				Choices: []Choice{
					{
						Text:         "\n\nThis is indeed a test",
						Index:        0,
						LogProbs:     nil,
						FinishReason: "length",
					},
				},
				Usage: Usage{
					PromptTokens:     5,
					CompletionTokens: 7,
					TotalTokens:      12,
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

			cReq := &CompletionRequest{
				Model:       TextDavinci003,
				Prompt:      "Say this is a test",
				MaxTokens:   7,
				Temperature: 0,
				TopP:        1,
				N:           1,
				Stream:      false,
				LogProbs:    nil,
				Stop:        "\n",
			}
			completionResp, err := c.RequestCompletion(context.TODO(), cReq)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expResp, completionResp)
		})
	}
}
