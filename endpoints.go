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

// IntelDetails endpoint.
type IntelDetails interface{}

// EntityType endpoint.
type EntityType interface{}

// Snapshot endpoint.
type Snapshot interface{}
