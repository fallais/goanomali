package goanomali

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

type IntelligenceResponse struct {
	Meta    Meta     `json:"meta"`
	Objects []Object `json:"objects"`
}

type Meta struct {
	Limit      int         `json:"limit"`
	Next       string      `json:"next"`
	Offset     int         `json:"offset"`
	Previous   interface{} `json:"previous"`
	TotalCount int         `json:"total_count"`
}

type Object struct {
	Itype                    string        `json:"itype"`
	ExpirationTs             time.Time     `json:"expiration_ts"`
	IP                       string        `json:"ip"`
	IsEditable               bool          `json:"is_editable"`
	FeedID                   int           `json:"feed_id"`
	UpdateID                 int64         `json:"update_id"`
	Value                    string        `json:"value"`
	IsPublic                 bool          `json:"is_public"`
	ThreatType               string        `json:"threat_type"`
	Workgroups               []interface{} `json:"workgroups"`
	Rdns                     interface{}   `json:"rdns"`
	Confidence               int           `json:"confidence"`
	UUID                     string        `json:"uuid"`
	TrustedCircleIds         []int         `json:"trusted_circle_ids"`
	ID                       int64         `json:"id"`
	Source                   string        `json:"source"`
	OwnerOrganizationID      int           `json:"owner_organization_id"`
	ImportSessionID          interface{}   `json:"import_session_id"`
	Latitude                 float64       `json:"latitude"`
	Type                     string        `json:"type"`
	Status                   string        `json:"status"`
	Description              interface{}   `json:"description"`
	Tags                     []Tags        `json:"tags"`
	Threatscore              int           `json:"threatscore"`
	SourceReportedConfidence int           `json:"source_reported_confidence"`
	ModifiedTs               time.Time     `json:"modified_ts"`
	Org                      string        `json:"org"`
	Asn                      string        `json:"asn"`
	CreatedTs                time.Time     `json:"created_ts"`
	Tlp                      interface{}   `json:"tlp"`
	IsAnonymous              interface{}   `json:"is_anonymous"`
	Country                  string        `json:"country"`
	Longitude                float64       `json:"longitude"`
	RetinaConfidence         int           `json:"retina_confidence"`
	ResourceURI              string        `json:"resource_uri"`
}

type Tags struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListIntelligence returns the indicators with given cirterias.
func (endpoint *Endpoint) ListIntelligence(ctx context.Context, query string, perPage, offset int) (*IntelligenceResponse, error) {
	// Prepare the URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/v2/intelligence"
	parameters := url.Values{}
	parameters.Add("username", endpoint.client.Username)
	parameters.Add("api_key", endpoint.client.APIKey)
	parameters.Add("limit", strconv.Itoa(perPage))
	parameters.Add("offset", strconv.Itoa(offset))
	parameters.Add("q", query)
	reqURL.RawQuery = parameters.Encode()

	// Create the request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}
	req = req.WithContext(ctx)

	// Set HTTP headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing the request: %s", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request: %s", err)
	}

	// Unmarshal the response
	var response *IntelligenceResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling the response : %s", err)
	}

	return response, nil
}
