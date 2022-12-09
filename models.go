package ggpt3

import (
	"context"
	"fmt"
	"net/http"
)

// Reference: https://beta.openai.com/docs/models/overview

// The OpenAI API is powered by a family of models with different capabilities and price points.
// You can also customize our base models for your specific use case with fine-tuning.

// GPT: A set of models that can understand and generate natural language
// Codex: A set of models that can understand and generate code, including translating natural language to code
// Content Filter: 	A fine-tuned model that can detect whether text may be sensitive or unsafe

const (
	// Davinci is the most capable model family and can perform any task the other models can
	// perform and often with less instruction. For applications requiring a lot of
	// understanding of the content, like summarization for a specific audience and
	// creative content generation, Davinci is going to produce the best results.
	// These increased capabilities require more compute resources, so Davinci costs
	// more per API call and is not as fast as the other models.
	//
	// Good at: Complex intent, cause and effect, summarization for audience
	TextDavinci003 = "text-davinci-003"

	// Curie is extremely powerful, yet very fast. While Davinci is stronger when it
	// comes to analyzing complicated text, Curie is quite capable for many nuanced
	// tasks like sentiment classification and summarization. Curie is also quite
	// good at answering questions and performing Q&A and as a general service chatbot.
	//
	// Good at: Language translation, complex classification, text sentiment, summarization
	TextCurie001 = "text-curie-001"

	// Babbage can perform straightforward tasks like simple classification. It’s also quite
	// capable when it comes to Semantic Search ranking how well documents match up with search queries.
	//
	// Good at: Moderate classification, semantic search classification
	TextBabbage001 = "text-babbage-001"

	// Ada is usually the fastest model and can perform tasks like parsing text, address correction
	// and certain kinds of classification tasks that don’t require too much nuance.
	// Ada’s performance can often be improved by providing more context.
	//
	// Good at: Parsing text, simple classification, address correction, keywords
	TextAda001 = "text-ada-001"

	// Most capable Codex model. Particularly good at translating natural language to code.
	// In addition to completing code, also supports inserting completions within code.
	CodexCodeDavinci002 = "code-davinci-002"

	// Almost as capable as Davinci Codex, but slightly faster.
	// This speed advantage may make it preferable for real-time applications.
	CodexCodeCushman001 = "code-cushman-001"

	TextDavinciEdit001 = "text-davinci-edit-001"
)

// While Davinci is generally the most capable, the other models can perform certain tasks
// extremely well with significant speed or cost advantages. For example, Curie can
// perform many of the same tasks as Davinci, but faster and for 1/10th the cost.

// The main GPT-3 models are meant to be used with the text completion endpoint.

// Models lists the currently available models, and provides basic information about each one such as the owner and availability.
func (c *Client) Models(ctx context.Context) (*ModelResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, nil, "/models", nil)
	if err != nil {
		return nil, err
	}

	var modelResp ModelResponse
	err = c.sendRequest(req, &modelResp)

	return &modelResp, err
}

// Model retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *Client) Model(ctx context.Context, model string) (*ModelResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, nil, fmt.Sprintf("/models/%s", model), nil)
	if err != nil {
		return nil, err
	}

	var modelResp ModelResponse
	err = c.sendRequest(req, &modelResp)

	return &modelResp, err
}
