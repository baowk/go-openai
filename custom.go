package openai

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateChatCompletionCustom(
	ctx context.Context,
	request ChatCompletionRequest, urlSuffix string,
) (response ChatCompletionResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	// urlSuffix := chatCompletionsSuffix
	// if !checkEndpointSupportsModel(urlSuffix, request.Model) {
	// 	err = ErrChatCompletionInvalidModel
	// 	return
	// }

	// if err = validateRequestForO1Models(request); err != nil {
	// 	return
	// }

	req, err := c.newRequest(
		ctx,
		http.MethodPost,
		c.fullURLSimple(urlSuffix),
		withBody(request),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) CreateCompletionCustom(ctx context.Context,
	request CompletionRequest, urlSuffix string) (response CompletionResponse, err error) {
	if request.Stream {
		err = ErrCompletionStreamNotSupported
		return
	}

	// if !checkEndpointSupportsModel(urlSuffix, request.Model) {
	// 	err = ErrCompletionUnsupportedModel
	// 	return
	// }

	// if !checkPromptType(request.Prompt) {
	// 	err = ErrCompletionRequestPromptTypeNotSupported
	// 	return
	// }

	req, err := c.newRequest(
		ctx,
		http.MethodPost,
		c.fullURLSimple(urlSuffix),
		withBody(request),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateChatCompletionStreamCustom(
	ctx context.Context,
	request ChatCompletionRequest, urlSuffix string,
) (stream *ChatCompletionStream, err error) {
	// urlSuffix := chatCompletionsSuffix
	// if !checkEndpointSupportsModel(urlSuffix, request.Model) {
	// 	err = ErrChatCompletionInvalidModel
	// 	return
	// }

	request.Stream = true
	if err = validateRequestForO1Models(request); err != nil {
		return
	}

	req, err := c.newRequest(
		ctx,
		http.MethodPost,
		c.fullURL(urlSuffix, withModel(request.Model)),
		withBody(request),
	)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[ChatCompletionStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &ChatCompletionStream{
		streamReader: resp,
	}
	return
}

func (c *Client) fullURLSimple(suffix string) string {
	baseURL := c.config.BaseURL
	if strings.HasPrefix(suffix, "/") {
		baseURL = strings.TrimRight(baseURL, "/")
	}
	return fmt.Sprintf("%s%s", baseURL, suffix)
}
