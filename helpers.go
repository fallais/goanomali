package goanomali

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) do(ctx context.Context, method, endpoint string, opts ...Option) (*http.Response, error) {
	// Options
	var apiOptions options

	// Add options
	for _, op := range opts {
		err := op(&apiOptions)
		if err != nil {
			return nil, err
		}
	}

	// Raw URL
	rawURL := fmt.Sprintf("%s/api/%s", c.BaseURL, endpoint)

	// Build query
	queryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// Assign query parameters
	if apiOptions.Params != nil {
		queryURL.RawQuery = apiOptions.Params.Encode()
	}

	// Initialize request
	req, err := http.NewRequest(method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add context
	req = req.WithContext(ctx)

	// Default headers
	headers := http.Header{}
	headers.Add("Accept", "application/json")

	// Optional headers
	if apiOptions.Headers != nil {
		for k := range *apiOptions.Headers {
			headers.Add(k, apiOptions.Headers.Get(k))
		}
	}

	// Assign new headers
	req.Header = headers

	// Do the query
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing the request: %s", err)
	}

	return resp, err
}
