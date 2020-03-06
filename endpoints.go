package goanomali

import (
	"context"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Endpoint is an API endpoint.
type Endpoint struct {
	client *Client
}

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// Intelligence endpoint.
type Intelligence interface {
	ListIntelligence(context.Context, string, int, int) (*IntelligenceResponse, error)
}
