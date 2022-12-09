package ggpt3

// ModelResponse represents a response from the Models v1 endpoint.
type ModelResponse struct {
	Data []data `json:"data"`
}

// CompletionResponse represents a reponse from the Completions v1 endpoint.
type CompletionResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

// EditsResponse represents a response from the edits endpoint.
type EditsResponse struct {
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

// ImageResponse represents a response from the images endpoint.
type ImageResponse struct {
	Created int
	Data    []data
}

type permission struct {
	Id                  string  `json:"id"`
	Object              string  `json:"object"`
	Created             int     `json:"created"`
	AllowCreateEngine   bool    `json:"allow_create_engine"`
	AllowSampling       bool    `json:"allow_sampling"`
	AllowLogProbs       bool    `json:"allow_logprobs"`
	AllowSearchIndicies bool    `json:"allow_search_indices"`
	AllowView           bool    `json:"allow_view"`
	AllowFineTuning     bool    `json:"allow_fine_tuning"`
	Organization        string  `json:"organization"`
	Group               *string `json:"group"`
	IsBlocking          bool    `json:"is_blocking"`
}

type data struct {
	Id         string       `json:"id"`
	Object     string       `json:"object"`
	OwnedBy    string       `json:"owned_by"`
	Permission []permission `json:"permission"`
	Url        string       `json:"url"`
}

type choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	LogProbs     *int   `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
