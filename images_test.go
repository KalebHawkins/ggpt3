package ggpt3

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	tcs := []struct {
		name     string
		apiKey   string
		iReq     *ImageRequest
		expIResp *ImageResponse
		expErr   error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			iReq: &ImageRequest{
				Prompt: "A test",
				N:      1,
			},
			expIResp: &ImageResponse{
				Created: 1589478378,
				Data: []data{
					{
						Url: "https://tisbutatest.com",
					},
				},
			},
			expErr: nil,
		},
		{
			name:   "WithApiErr",
			apiKey: "",
			expErr: errApiKeyUndefined,
		},
	}

	ts := startTestServer(t)
	defer ts.Close()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(tc.apiKey)
			c.baseUrl = ts.URL

			if tc.expErr == errApiKeyUndefined {
				c.apiKey = ""
			}

			imageResp, err := c.RequestImages(context.TODO(), tc.iReq)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expIResp, imageResp)
		})
	}
}

func TestImageEdits(t *testing.T) {
	tcs := []struct {
		name     string
		apiKey   string
		iReq     *ImageRequest
		expIResp *ImageResponse
		expErr   error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			iReq: &ImageRequest{
				Image:  "testdata/otter.png",
				Prompt: "A test",
				N:      1,
			},
			expIResp: &ImageResponse{
				Created: 1589478378,
				Data: []data{
					{
						Url: "https://tisbutatest.com",
					},
				},
			},
			expErr: nil,
		},
		{
			name:   "WithApiErr",
			apiKey: "",
			iReq: &ImageRequest{
				Image:  "testdata/otter.png",
				Prompt: "A Test",
				N:      1,
			},
			expErr: errApiKeyUndefined,
		},
		{
			name:   "ErrNoImage",
			apiKey: "",
			iReq: &ImageRequest{
				Prompt: "A Test",
				N:      1,
			},
			expErr: fmt.Errorf("image is a required field for an edits request"),
		},
		{
			name:   "ErrNoPrompt",
			apiKey: "",
			iReq: &ImageRequest{
				Image: "testdata/otter.png",
				N:     1,
			},
			expErr: fmt.Errorf("prompt is a required field for an edits request"),
		},
	}

	ts := startTestServer(t)
	defer ts.Close()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(tc.apiKey)
			c.baseUrl = ts.URL

			if tc.expErr == errApiKeyUndefined {
				c.apiKey = ""
			}

			imageResp, err := c.RequestImageEdits(context.TODO(), tc.iReq)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expIResp, imageResp)
		})
	}
}

func TestImageVariations(t *testing.T) {
	tcs := []struct {
		name     string
		apiKey   string
		iReq     *ImageRequest
		expIResp *ImageResponse
		expErr   error
	}{
		{
			name:   "WithoutErr",
			apiKey: testApiKey,
			iReq: &ImageRequest{
				Image:  "testdata/otter.png",
				Prompt: "A test",
				N:      1,
			},
			expIResp: &ImageResponse{
				Created: 1589478378,
				Data: []data{
					{
						Url: "https://tisbutatest.com",
					},
				},
			},
			expErr: nil,
		},
		{
			name:   "WithApiErr",
			apiKey: "",
			iReq: &ImageRequest{
				Image:  "testdata/otter.png",
				Prompt: "A Test",
				N:      1,
			},
			expErr: errApiKeyUndefined,
		},
		{
			name:   "ErrNoImage",
			apiKey: "",
			iReq: &ImageRequest{
				Prompt: "A Test",
				N:      1,
			},
			expErr: fmt.Errorf("image is a required field for a variation request"),
		},
	}

	ts := startTestServer(t)
	defer ts.Close()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(tc.apiKey)
			c.baseUrl = ts.URL

			if tc.expErr == errApiKeyUndefined {
				c.apiKey = ""
			}

			imageResp, err := c.RequestImageVariations(context.TODO(), tc.iReq)

			if tc.expErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expErr, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expIResp, imageResp)
		})
	}
}
