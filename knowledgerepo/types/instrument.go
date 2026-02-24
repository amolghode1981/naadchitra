package types

/*
*
Currently this code does not consider the possibility of multiple performers playing the same instrument in a painting.
For the first draft it is ok.
*/
type InstrumentType string

const (
	InstrumentClassChordophones   = "tanta"
	InstrumentClassAerophones     = "sushira"
	InstrumentClassMembranophones = "avanaddha"
	InstrumentClassIdiophone      = "ghana"
)

type InstrumentName string

const (
	InstrumentPakhawaj = "pakhawaj"
	InstrumentTabla    = "tabla"
	InstrumentSitar    = "sitar"
	InstrumentSarangi  = "sarangi"
	InstrumentDholak   = "dholak"
	InstrumentManjira  = "manjira"
	InstrumentBeen     = "been"
	InstrumentRabab    = "rabab"
	InstrumentTambora  = "tambora"
)

/*
This is the index mapping of the instruments in the CSV. It is not yet foolproof.
Nonetheless, it is a good starting point for parsing the performers and their instruments from the CSV.
TODO(Amol): Unify the way
*/
type InstrumentSequence int

const (
	Pakhawaj             = 0
	Manjira              = 2
	Been                 = 4
	Rabab                = 6
	Tambora              = 8
	Singer               = 10
	Sitar                = 12
	Sarangi              = 14
	Dholak               = 16
	Tabla                = 18
	Additional           = 20
	TaalIndicatingPerson = 21
)

type PlayingTechnique string

const (
	PlayingTechniqueFingerPlucked        = "finger_plucked"
	PlayingTechniqueBowed                = "bowing"
	PlayingTechniqueHandStruck           = "hand_struck"
	PlayingTechniqueMalletStruck         = "mallet_struck"
	PlayingTechniqueObjectToObjectStruck = "object_to_object_struck"
	PlayingTechniqueBlown                = "blown"
)

type Instrument struct {
	Name             InstrumentName   `json:"name"`
	Class            InstrumentType   `json:"class"`
	PlayingTechnique PlayingTechnique `json:"playing_technique"`
}

func (m *Instrument) CreatePakhawaj() (Instrument, error) {
	m.Class = InstrumentClassMembranophones
	m.Name = InstrumentPakhawaj
	m.PlayingTechnique = PlayingTechniqueHandStruck
	return *m, nil
}

func (m *Instrument) CreateManjira() (Instrument, error) {
	m.Class = InstrumentClassIdiophone
	m.Name = InstrumentManjira
	m.PlayingTechnique = PlayingTechniqueObjectToObjectStruck
	return *m, nil
}

func (m *Instrument) CreateBeen() (Instrument, error) {
	m.Class = InstrumentClassChordophones
	m.Name = InstrumentBeen
	m.PlayingTechnique = PlayingTechniqueFingerPlucked
	return *m, nil
}

func (m *Instrument) CreateRabab() (Instrument, error) {
	m.Class = InstrumentClassChordophones
	m.Name = InstrumentRabab
	m.PlayingTechnique = PlayingTechniqueFingerPlucked
	return *m, nil
}

func (m *Instrument) CreateTambora() (Instrument, error) {
	m.Class = InstrumentClassChordophones
	m.Name = InstrumentTambora
	m.PlayingTechnique = PlayingTechniqueFingerPlucked
	return *m, nil
}

func (m *Instrument) CreateSitar() (Instrument, error) {
	m.Class = InstrumentClassChordophones
	m.Name = InstrumentSitar
	m.PlayingTechnique = PlayingTechniqueFingerPlucked
	return *m, nil
}

func (m *Instrument) CreateSarangi() (Instrument, error) {
	m.Class = InstrumentClassChordophones
	m.Name = InstrumentSarangi
	m.PlayingTechnique = PlayingTechniqueBowed
	return *m, nil
}

func (m *Instrument) CreateDholak() (Instrument, error) {
	m.Class = InstrumentClassMembranophones
	m.Name = InstrumentDholak
	m.PlayingTechnique = PlayingTechniqueHandStruck
	return *m, nil
}

func (m *Instrument) CreateTabla() (Instrument, error) {
	m.Class = InstrumentClassMembranophones
	m.Name = InstrumentTabla
	m.PlayingTechnique = PlayingTechniqueHandStruck
	return *m, nil
}
