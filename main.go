package main

import (
	"NAADCHITRA/connectors"
	knowledgerepo "NAADCHITRA/knowledgerepo/import"
	"NAADCHITRA/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Registry maps museum keys to their connectors
var registry = map[string]connectors.Connector{
	"met": connectors.NewMetConnector(),
}

// corsMiddleware adds CORS headers to every response and handles preflight.
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Vite dev server
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Preflight request — browser sends this before the real request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func objectHandler(w http.ResponseWriter, r *http.Request) {
	museum := r.URL.Query().Get("museum")
	id := r.URL.Query().Get("id")

	if museum == "" || id == "" {
		http.Error(w, "missing 'museum' or 'id' query params", http.StatusBadRequest)
		return
	}

	connector, ok := registry[museum]
	if !ok {
		http.Error(w, fmt.Sprintf("unknown museum: %s", museum), http.StatusNotFound)
		return
	}

	obj, err := connector.FetchObject(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("fetch error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func museumsHandler(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keys)
}

// func main() {
// 	http.HandleFunc("/api/object", corsMiddleware(objectHandler))
// 	http.HandleFunc("/api/museums", corsMiddleware(museumsHandler))

// 	log.Println("Backend running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	paintings, err := knowledgerepo.ImportCSV("knowledgerepo/import/csv/met_550e8400-e29b-41d4-a716-446655440000.csv")
	if err != nil {
		log.Error().Err(err).Msg("Failed to import CSV")
		return
	}

	data, error := utils.ToPrettyJSON(paintings)
	if error != nil {
		log.Error().Err(error).Msg("Failed to convert paintings to JSON")
		return
	}
	log.Info().Msgf("Parsed paintings: count=%d, data=%s", len(paintings), data)
}
