package connectors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const metBaseURL = "https://collectionapi.metmuseum.org/public/collection/v1"

// metResponse mirrors the Met API response shape.
type metResponse struct {
	ObjectID          int    `json:"objectID"`
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	Medium            string `json:"medium"`
	Dimensions        string `json:"dimensions"`
	Department        string `json:"department"`
	Culture           string `json:"culture"`
	PrimaryImage      string `json:"primaryImage"`
	ObjectURL         string `json:"objectURL"`
	Tags              []struct {
		Term string `json:"term"`
	} `json:"tags"`
}

// MetConnector fetches objects from the Metropolitan Museum of Art.
type MetConnector struct {
	client *http.Client
}

func NewMetConnector() *MetConnector {
	return &MetConnector{client: &http.Client{}}
}

func (m *MetConnector) Name() string {
	return "The Metropolitan Museum of Art"
}

func (m *MetConnector) FetchObject(id string) (*ArtObject, error) {
	url := fmt.Sprintf("%s/objects/%s", metBaseURL, id)

	resp, err := m.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Met API returned status %d", resp.StatusCode)
	}

	var raw metResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Normalize tags
	tags := make([]string, 0, len(raw.Tags))
	for _, t := range raw.Tags {
		tags = append(tags, t.Term)
	}

	return &ArtObject{
		Museum:     m.Name(),
		ID:         fmt.Sprintf("%d", raw.ObjectID),
		Title:      raw.Title,
		Artist:     raw.ArtistDisplayName,
		Date:       raw.ObjectDate,
		Medium:     raw.Medium,
		Dimensions: raw.Dimensions,
		Department: raw.Department,
		Culture:    raw.Culture,
		ImageURL:   raw.PrimaryImage,
		SourceURL:  raw.ObjectURL,
		Tags:       tags,
	}, nil
}
