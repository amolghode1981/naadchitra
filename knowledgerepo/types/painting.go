package types

import (
	"fmt"
	"path"
	"strconv"

	"github.com/google/uuid"
)

type Painting struct {
	ObjectId         string             `json:"objectID"`
	MuseumObjectId   string             `json:"museumObjectID"`
	URL              string             `json:"url,omitempty"`
	ImageUrl         string             `json:"url,omitempty"`
	Museum           string             `json:"museum"`
	Dynasty          string             `json:"dynasty"`
	Date             string             `json:"date"`
	Performers       []Performer        `json:"performers,omitempty"`
	Painter          string             `json:"painter,omitempty"`
	TypeOfPainting   PaintingTypeString `json:"typeOfPainting,omitempty"`
	NameOfThePainter string             `json:"nameOfThePainter,omitempty"`
	Notes            string             `json:"notes,omitempty"`
	/* These two Ids are crucian and the represent the core knowledge
	of the painting and music making. They link to the musical and cultural semantic notes respectively.
	They are the object ids of the records created by the experts in the field*/
	MusicalSematicNotesId   string `json:"musicalSemanticNotesId,omitempty"`
	CulturalSemanticNotesId string `json:"culturalSemanticNotesId,omitempty"`
	/* If the musician creating this information decides to add a audio-visual reference, this list
	represents the objectIds of those references
	*/
	AudioVisualReferenceNotesIds []string `json:"audioVisualReferenceNotesIds"`
}

// Row Collection,Source,Approx Date,Dynesty,Pakhawaj,Taal indicating person,Manjira,Been,Rabab,Tambur/Tambora,Singer,Sitar,Sarangi,Dholak,Tabla,Additional instruments,Gender of the performers,Type of painting,Name of the artist if available,Comments
func (m *Painting) ParseRow(row []string) (Painting, error) {
	fmt.Println(row)
	m.ObjectId = uuid.New().String()
	m.Museum = row[0]
	objectId := ""
	if len(row[1]) == 0 {
		objectId = path.Base(row[2])
	}
	m.MuseumObjectId = objectId
	m.URL = row[2]
	m.Date = row[3]
	m.Dynasty = row[4]
	m.ImageUrl = ""

	// Slice from 5 to 28 for the performers and their instruments.
	// 28 is not included.
	performers, err := ParsePerformers(row[5:28])
	if err != nil {
		return *m, err
	}
	m.Performers = performers
	// gender of the performer columl : Obsolete, row[28]
	// Type of painting row[29]	Name of the artist if available [30]	Comments [31]

	n, err := strconv.Atoi(row[29])
	if err == nil {
		m.TypeOfPainting = PaintingTypeByInt[n].StringValue
	}
	m.NameOfThePainter = row[30]
	m.Notes = row[31]
	fmt.Println("returning")
	return *m, nil
}
