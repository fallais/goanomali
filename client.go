package goanomali

import (
	"net/http"
)

const (
	defaultVersion = "12.0"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Client is a client for QRadar REST API.
type Client struct {
	client *http.Client

	// BaseURL is the base URL for API requests.
	BaseURL string

	// Username is the username.
	Username string

	// APIKey is the API key.
	APIKey string

	// Endpoints
	Intelligence Intelligence
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewClient returns a new QRadar API client.
func NewClient(httpClient *http.Client, baseURL, username, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Create the client
	c := &Client{
		client:   httpClient,
		BaseURL:  baseURL,
		Username: username,
		APIKey:   apiKey,
	}

	// Add the endpoints
	c.Intelligence = &Endpoint{client: c}

	return c
}
