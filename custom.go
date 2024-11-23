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

func (c *Client) fullURLSimple(suffix string) string {
	baseURL := c.config.BaseURL
	if strings.HasPrefix(suffix, "/") {
		baseURL = strings.TrimRight(baseURL, "/")
	}
	return fmt.Sprintf("%s%s", baseURL, suffix)
}
