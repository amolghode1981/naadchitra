package knowledgerepo

/*
Cant import individual types. The entire package is imported
*/
import (
	"NAADCHITRA/knowledgerepo/types"
	"encoding/csv"
	"fmt"
	"os"
)

func ImportCSV(filePath string) ([]types.Painting, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all rows at once
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var paintings []types.Painting
	// Loop through rows, skipping the header
	for _, row := range rows[1:] {
		var painting types.Painting
		painting, err := painting.ParseRow(row)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row: %w", err)
		}
		// painting := types.Painting{
		// 	ObjectId:                     parseInt(row[0]),
		// 	MuseumObjectId:               row[1],
		// 	Museum:                       row[2],
		// 	Dynasty:                      row[3],
		// 	Date:                         row[4],
		// 	Painter:                      row[5],
		// 	MusicalSematicNotesId:        row[6],
		// 	CulturalSemanticNotesId:      row[7],
		// 	AudioVisualReferenceNotesIds: parseAudioVisualReferences(row[8]),
		// }
		paintings = append(paintings, painting)
	}

	return paintings, nil
}
