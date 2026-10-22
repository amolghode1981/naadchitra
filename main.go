package main

import (
	"NAADCHITRA/connectors"
	"io"
	"os"
	"path"

	// pb "NAADCHITRA/prediction_rpc"

	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Registry maps museum keys to their connectors
var registry = map[string]connectors.Connector{
	"met":       connectors.NewMetConnector(),
	"cleveland": connectors.NewClevelandConnector(),
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

// 	// log.Println("Backend running at http://localhost:8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal().Err(err).Msg("Server failed")
// 	}
// }

func extractLastPart(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return path.Base(u.Path)
}

func main() {
	//inFile, err := os.Open("knowledgerepo/import/csv/cleveland_550e8400-e29b-41d4-a716-446655440000_back.csv")
	inFile, err := os.Open("knowledgerepo/import/csv/cleveland_base.csv")
	// inFile, err := os.Open("knowledgerepo/import/csv/met_550e8400-e29b-41d4-a716-446655440000.csv")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create("knowledgerepo/import/csv/cleveland_output.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	reader := csv.NewReader(inFile)
	reader.FieldsPerRecord = -1
	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Ensure at least 3 columns
		if len(record) < 3 {
			continue
		}

		// Step 1: Replace column 1 with extracted value from URL (column 2)
		if record[1] == "" && record[2] != "" {
			record[1] = extractLastPart(record[2])
		}

		// Step 2: Replace empty fields with "3"
		for i := range record {
			if record[i] == "" {
				record[i] = "3"
			}
		}

		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
}

// func main() {
// 	// paintings, err := knowledgerepo.ImportCSV("knowledgerepo/import/csv/met_550e8400-e29b-41d4-a716-446655440000.csv")
// 	// paintings, err := knowledgerepo.ImportCSV("knowledgerepo/import/csv/cleveland_550e8400-e29b-41d4-a716-446655440000.csv")
// 	paintings, err := knowledgerepo.ImportCSV("knowledgerepo/import/csv/cleveland_base.csv")
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to import CSV")
// 		return
// 	}

// 	for i := 0; i < len(paintings); i++ {
// 		connector, ok := registry["cleveland"]
// 		if !ok {
// 			log.Info().Msg("DEBUG_1")
// 			log.Error().Msg("Failed to import CSV")
// 			return
// 		}
// 		museumData, err := connector.FetchObject(paintings[i].MuseumObjectId)
// 		if err != nil {
// 			log.Info().Msg("DEBUG_2")
// 			log.Info().Msg(paintings[i].MuseumObjectId)
// 			log.Error().Err(err).Msg("Failed to import CSV")
// 			return
// 		}
// 		paintings[i].ImageUrl = museumData.ImageURL
// 		log.Info().Msg(fmt.Sprintf("ObjectId = %s url = %s", museumData.ID, museumData.ImageURL))
// 	}

// 	// data, error := utils.ToPrettyJSON(paintings)
// 	// if error != nil {
// 	// 	log.Error().Err(error).Msg("Failed to convert paintings to JSON")
// 	// 	return
// 	// }
// 	// log.Info().Msgf("Parsed paintings: count=%d, data=%s", len(paintings), data)
// }

// func main() {
// 	// Connect to server
// 	conn, err := grpc.Dial(
// 		"localhost:50051",
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("Failed to connect")
// 	}
// 	// Run this when function exits.
// 	defer conn.Close()

// 	// Create client stub
// 	client := pb.NewAskKosambiClient(conn)

// 	// Set timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// // Call guessPaintingStyle
// 	// resp, err := client.GuessPaintingStyle(
// 	// 	ctx,
// 	// 	&pb.StylePredictionRequest{
// 	// 		PaintingUrl: "https://www.metmuseum.org/art/collection/search/50939",
// 	// 	},
// 	// )
// 	// if err != nil {
// 	// 	log.Fatal().Err(err).Msg("Failed to call")
// 	// }

// 	// for _, style := range resp.Styles {
// 	// 	log.Printf("Style: %s Confidence: %d", style.Style, style.Confidence)
// 	// }

// 	// Call guessPaintingStyle
// 	resp, err := client.GuessMusicMaking(
// 		ctx,
// 		&pb.MusicPredictionRequest{
// 			// predicted yes expected yes PaintingUrl: "https://images.metmuseum.org/CRDImages/as/original/DP153182.jpg",
// 			// predicted no expected no PaintingUrl: "https://images.metmuseum.org/CRDImages/as/original/DP112943.jpg",
// 			// predicted yes expected yes  PaintingUrl: "https://images.metmuseum.org/CRDImages/as/original/DT4984.jpg",
// 			// predicted yes expected no PaintingUrl: "https://images.metmuseum.org/CRDImages/as/original/DP152284.jpg",
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("Failed to call")
// 	}

// 	log.Debug().Msgf("MusicMakingDepicted: %v Confidence: %d", resp.MusicMakingDepicted, resp.Confidence)
// }
