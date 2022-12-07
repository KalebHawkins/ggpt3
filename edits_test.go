package ggpt3

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestEdits(t *testing.T) {
	tcs := []struct {
		name    string
		apiKey  string
		expResp *EditsResponse
		expErr  error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			expResp: &EditsResponse{
				Object:  "edit",
				Created: 1589478378,
				Choices: []choice{{
					Text:  "What day of the week is it?",
					Index: 0,
				}},
				Usage: usage{
					PromptTokens:     25,
					CompletionTokens: 32,
					TotalTokens:      57,
				},
			},
			expErr: nil,
		},
		{
			name:    "WithErr",
			apiKey:  testApiKey,
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

			er := &EditsRequest{
				Model:       TextDavinciEdit001,
				Input:       "What day of the wek is it?",
				Instruction: "Fix the spelling mistakes",
			}
			edits, err := c.RequestEdits(context.TODO(), er)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expResp, edits)

		})
	}
}
