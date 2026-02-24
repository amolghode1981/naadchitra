package types

import "fmt"

type PerformerType string

const (
	PerformerIndicatingTaala = "indicating_taala"
	PerformerDancer          = "dancer"
	PerformerSinger          = "singer"
	PerformerInstrumentalist = "instrumentalist"
)

type Gender string

const (
	PerformerGenderMale   = "male"
	PerformerGenderFemale = "female"
	PerformerGenderOther  = "other"
)

type Performer struct {
	Gender               Gender        `json:"gender"`
	Type                 PerformerType `json:"type"`
	SpeculativeEthnicity string        `json:"speculative_ethnicity,omitempty"`
	AdditionalNotes      string        `json:"additionalNotes,omitempty"`
	Instrument           Instrument    `json:"instrument,omitempty"` // Only applicable for instrumentalists
}

func createPakhawajPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreatePakhawaj()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createManjiraPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateManjira()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createBeenPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateBeen()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createRababPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateRabab()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createTamboraPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateTambora()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createSinger(gender Gender) (Performer, error) {
	return Performer{
		Gender:     gender,
		Type:       PerformerSinger,
		Instrument: Instrument{},
	}, nil
}

func createSitarPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateSitar()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createSarangiPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateSarangi()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createDholakPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateDholak()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

func createTablaPerformer(gender Gender) (Performer, error) {
	var instrument, err = (&Instrument{}).CreateTabla()
	if err != nil {
		return Performer{}, err
	}
	return Performer{
		Gender:     gender,
		Type:       PerformerInstrumentalist,
		Instrument: instrument,
	}, nil
}

/*
*
Currently we are not parsing additional instrumnts as a part of the CSV import, but if we decide to add that in the future, this function can be used to create performers for those additional instruments
*/
func createAdditionalPerformer() (Performer, error) {
	return Performer{
		Gender:     PerformerGenderOther,
		Type:       PerformerInstrumentalist,
		Instrument: Instrument{},
	}, nil
}

func createTaalIndicatingPerson(gender Gender) (Performer, error) {
	return Performer{
		Gender:     gender,
		Type:       PerformerIndicatingTaala,
		Instrument: Instrument{},
	}, nil
}

func getGenderFromColumn(value string) (Gender, error) {
	switch value {
	case "0":
		return PerformerGenderFemale, nil
	case "1":
		return PerformerGenderMale, nil
	case "2":
		return PerformerGenderOther, nil
	default:
		return "", fmt.Errorf("invalid gender value: %s", value)
	}
}

func ParsePerformers(instruments []string) ([]Performer, error) {
	var performers []Performer
	var index int = 0
	var done bool = false
	for index < len(instruments) {
		var performer Performer
		var performerError error
		gender, genderError := getGenderFromColumn(instruments[index+1])
		switch index {
		case Pakhawaj:
			if instruments[index] != "" {
				performer, performerError = createPakhawajPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Manjira:
			if instruments[index] != "" {
				performer, performerError = createManjiraPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Been:
			if instruments[index] != "" {
				performer, performerError = createBeenPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Rabab:
			if instruments[index] != "" {
				performer, performerError = createRababPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Tambora:
			if instruments[index] != "" {
				performer, performerError = createTamboraPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Singer:
			if instruments[index] != "" {
				performer, performerError = createSinger(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Sitar:
			if instruments[index] != "" {
				performer, performerError = createSitarPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Sarangi:
			if instruments[index] != "" {
				performer, performerError = createSarangiPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Dholak:
			if instruments[index] != "" {
				performer, performerError = createDholakPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Tabla:
			if instruments[index] != "" {
				performer, performerError = createTablaPerformer(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 2
		case Additional:
			if instruments[index] != "" {
				performer, performerError = createAdditionalPerformer()
				if performerError == nil {
					performers = append(performers, performer)
				}
			}
			index += 1
		case TaalIndicatingPerson:
			if instruments[index] != "" {
				performer, performerError = createTaalIndicatingPerson(gender)
				if genderError == nil && performerError == nil {
					performers = append(performers, performer)
				}
			}
			done = true
		default:
			// Handle other instruments or skip
		}
		fmt.Println("Performer index", index, ":", performer)
		if done {
			break
		}
	}
	fmt.Println("Performers ", performers)
	return performers, nil
}
