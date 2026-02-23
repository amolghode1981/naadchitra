package connectors

// ArtObject is the normalized representation all museum connectors map to.
type ArtObject struct {
	Museum     string   `json:"museum"`
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Artist     string   `json:"artist"`
	Date       string   `json:"date"`
	Medium     string   `json:"medium"`
	Dimensions string   `json:"dimensions"`
	Department string   `json:"department"`
	Culture    string   `json:"culture"`
	ImageURL   string   `json:"imageUrl"`
	SourceURL  string   `json:"sourceUrl"`
	Tags       []string `json:"tags,omitempty"`
}

// Connector is the interface every museum adapter must implement.
type Connector interface {
	// Name returns a human-readable museum name.
	Name() string
	// FetchObject retrieves and normalizes an object by its museum-specific ID.
	FetchObject(id string) (*ArtObject, error)
}
