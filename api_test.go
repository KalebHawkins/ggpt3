package ggpt3

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testApiKey = "timeToGetTestyInHere"
)

var (
	errApiKeyUndefined = errors.New("error, status code 401, message: You didn't provide an API key")
)

func TestSendRequest(t *testing.T) {
	ts := startTestServer(t)
	defer ts.Close()

	expectedResponse := ErrorResponse{
		Error: &Error{
			Message: "Invalid URL (GET /v1)",
			Type:    "invalid_request_error",
			Param:   nil,
			Code:    nil,
		},
	}

	c := NewClient(testApiKey)
	req, err := http.NewRequest(http.MethodGet, ts.URL+"/v1", nil)
	assert.NoError(t, err)

	resp, err := c.httpClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var jsonResp ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, jsonResp)
}

func startTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1":
			mockV1Response(t, w, r)
		case "/models", "/models/model-id-0":
			mockV1ModelsResponse(t, w, r)
		case "/completions":
			mockV1CompletionResponse(t, w, r)
		case "/edits":
			mockV1EditsRequest(t, w, r)
		}
	}))
}

func mockV1Response(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	expectedResponse := ErrorResponse{
		Error: &Error{
			Message: "Invalid URL (GET /v1)",
			Type:    "invalid_request_error",
			Param:   nil,
			Code:    nil,
		},
	}

	w.WriteHeader(http.StatusNotFound)
	jsonResp, err := json.Marshal(expectedResponse)
	assert.NoError(t, err)
	w.Write(jsonResp)
}

func mockV1ModelsResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	if mockV1ApiKeyUndefined(t, w, r) {
		return
	}

	switch r.URL.Path {
	case "/models":
		resp, err := os.ReadFile("testdata/models.json")
		assert.NoError(t, err)
		w.Write(resp)
	case "/models/model-id-0":
		resp, err := os.ReadFile("testdata/model.json")
		assert.NoError(t, err)
		w.Write(resp)
	}
}

func mockV1CompletionResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	if mockV1ApiKeyUndefined(t, w, r) {
		return
	}

	resp, err := os.ReadFile("testdata/completionResponse.json")
	assert.NoError(t, err)
	w.Write(resp)
}

func mockV1ApiKeyUndefined(t *testing.T, w http.ResponseWriter, r *http.Request) bool {
	t.Helper()

	authKey := r.Header.Get("Authorization")

	if authKey != fmt.Sprintf("Bearer %s", testApiKey) {
		resp, err := os.ReadFile("testdata/error.json")
		assert.NoError(t, err)

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(resp)
		return true
	}

	return false
}

func mockV1EditsRequest(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	if mockV1ApiKeyUndefined(t, w, r) {
		return
	}

	resp, err := os.ReadFile("testdata/editsResponse.json")
	assert.NoError(t, err)
	w.Write(resp)
}
