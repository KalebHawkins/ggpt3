package ggpt3

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	tcs := []struct {
		name        string
		expResponse *ModelResponse
		expErr      error
	}{
		{
			name: "WithoutErr",
			expResponse: &ModelResponse{
				Data: []Data{
					{
						Id:      "model-id-0",
						Object:  "model",
						OwnedBy: "organization-owner",
					},
					{
						Id:      "model-id-1",
						Object:  "model",
						OwnedBy: "organization-owner",
					},
				},
			},
			expErr: nil,
		},
		{
			name:        "WithErr",
			expResponse: nil,
			expErr:      errApiKeyUndefined,
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

			models, err := c.Models(context.TODO())

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expResponse, models)
		})
	}
}

func TestModel(t *testing.T) {
	tcs := []struct {
		name        string
		expResponse *ModelResponse
		expErr      error
	}{
		{
			name: "WithoutErr",
			expResponse: &ModelResponse{
				Data: []Data{
					{
						Id:      "model-id-0",
						Object:  "model",
						OwnedBy: "organization-owner",
					},
				},
			},
			expErr: nil,
		},
		{
			name:        "WithErr",
			expResponse: nil,
			expErr:      errApiKeyUndefined,
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

			models, err := c.Model(context.TODO(), "model-id-0")

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expResponse, models)
		})
	}
}
