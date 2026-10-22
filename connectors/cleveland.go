package connectors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const clevelandBaseURL = "https://openaccess-api.clevelandart.org/api/artworks/"

// clevelandResponse mirrors the Met API response shape.
type Response struct {
	Data clevelandResponse `json:"data"`
}
type clevelandResponse struct {
	ID              int    `json:"id"`
	AccessionNumber string `json:"accession_number"`
	// ShareLicenseStatus    string            `json:"share_license_status"`
	// Tombstone             string            `json:"tombstone"`
	// CurrentLocation       *string           `json:"current_location"`
	Title string `json:"title"`
	// CreationDate          string            `json:"creation_date"`
	// CreationDateEarliest  int               `json:"creation_date_earliest"`
	// CreationDateLatest    int               `json:"creation_date_latest"`
	// ArtistsTags           []string          `json:"artists_tags"`
	// Culture               []string          `json:"culture"`
	// Technique             string            `json:"technique"`
	// SupportMaterials      []string          `json:"support_materials"`
	// Department            string            `json:"department"`
	// Collection            string            `json:"collection"`
	// Type                  string            `json:"type"`
	// Measurements          string            `json:"measurements"`
	// Dimensions            Dimensions        `json:"dimensions"`
	// StateOfTheWork        *string           `json:"state_of_the_work"`
	// EditionOfTheWork      *string           `json:"edition_of_the_work"`
	// Copyright             *string           `json:"copyright"`
	// Inscriptions          []string          `json:"inscriptions"`
	// Exhibitions           Exhibitions       `json:"exhibitions"`
	// FindSpot              *string           `json:"find_spot"`
	// RelatedWorks          []string          `json:"related_works"`
	// FormerAccessionNumber []string          `json:"former_accession_numbers"`
	// DidYouKnow            *string           `json:"did_you_know"`
	// Description           *string           `json:"description"`
	// ExternalResources     ExternalResources `json:"external_resources"`
	// Citations             []string          `json:"citations"`
	// URL                   string            `json:"url"`
	Images Images `json:"images"`
	// AlternateImages       []AlternateImage  `json:"alternate_images"`
	// Creditline            string            `json:"creditline"`
	// ImageCredit           *string           `json:"image_credit"`
	// SketchfabID           *string           `json:"sketchfab_id"`
	// SketchfabURL          *string           `json:"sketchfab_url"`
	// GalleryDonorText      *string           `json:"gallery_donor_text"`
	// AthenaID              int               `json:"athena_id"`
	// Creators              []string          `json:"creators"`
	// LegalStatus           string            `json:"legal_status"`
	// AccessionDate         string            `json:"accession_date"`
	// SortableDate          int               `json:"sortable_date"`
	// DateAddedToOA         *string           `json:"date_added_to_oa"`
	// DateText              string            `json:"date_text"`
	// CollapseArtists       bool              `json:"collapse_artists"`
	// OnLoan                bool              `json:"on_loan"`
	// RecentlyAcquired      bool              `json:"recently_acquired"`
	// RecordType            string            `json:"record_type"`
	// ConservationStatement *string           `json:"conservation_statement"`
	// HasConservationImages bool              `json:"has_conservation_images"`
	// CoverAccessionNumber  *string           `json:"cover_accession_number"`
	// IsNaziEraProvenance   bool              `json:"is_nazi_era_provenance"`
	// Impression            *string           `json:"impression"`
	// AlternateTitles       []string          `json:"alternate_titles"`
	// IsHighlight           bool              `json:"is_highlight"`
	// UpdatedAt             string            `json:"updated_at"`
}

type Dimensions struct {
	Painting PaintingDimensions `json:"painting"`
}

type PaintingDimensions struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type Exhibitions struct {
	Current []interface{}      `json:"current"`
	Legacy  []LegacyExhibition `json:"legacy"`
}

type LegacyExhibition struct {
	Description string `json:"description"`
	OpeningDate string `json:"opening_date"`
}

type ExternalResources struct {
	Wikidata        []string `json:"wikidata"`
	InternetArchive []string `json:"internet_archive"`
}

type Images struct {
	Annotation *string   `json:"annotation"`
	Web        ImageInfo `json:"web"`
	Print      ImageInfo `json:"print"`
	Full       ImageInfo `json:"full"`
}

type AlternateImage struct {
	DateCreated string    `json:"date_created"`
	Annotation  string    `json:"annotation"`
	Web         ImageInfo `json:"web"`
	Print       ImageInfo `json:"print"`
	Full        ImageInfo `json:"full"`
}

type ImageInfo struct {
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Filesize string `json:"filesize"`
	Filename string `json:"filename,omitempty"`
}

// ClevelandConnector fetches objects from the Metropolitan Museum of Art.
type ClevelandConnector struct {
	client *http.Client
}

func NewClevelandConnector() *ClevelandConnector {
	return &ClevelandConnector{client: &http.Client{}}
}

func (m *ClevelandConnector) Name() string {
	return "The Cleveland Museum of Art"
}

func (m *ClevelandConnector) FetchObject(id string) (*ArtObject, error) {
	url := fmt.Sprintf("%s%s", clevelandBaseURL, id)
	print(url)
	resp, err := m.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cleveland API returned status %d", resp.StatusCode)
	}

	var raw Response
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &ArtObject{
		Museum:   m.Name(),
		ID:       fmt.Sprintf("%s", raw.Data.AccessionNumber),
		Title:    raw.Data.Title,
		ImageURL: raw.Data.Images.Web.URL,
	}, nil
}
