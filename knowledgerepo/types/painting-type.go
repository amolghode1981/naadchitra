package types

import "fmt"

// 1: Raagmala, 2: Court scene, 3 Common life, 4: war, 5: Scene with Gods but not Raagmala (Usually RaagDhyans in Raagmala are the deities), 6: Personal 7: Casual/Formal gathering which is neither a personal spaces, nor a common life. E.g musicians sitting together, no king, no chores of life, marriage
type PaintingType int

const (
	_                 = iota
	RaagMala          = 1
	CourtScene        = 2
	CommonLife        = 3
	War               = 4
	SceneWithGods     = 5
	Personal          = 6
	NonCourtGathering = 7
)

type PaintingTypeString string

const (
	TypeOfPaintingRaagMala          = "raag_mala"
	TypeOfPaintingCourtScene        = "court_scene"
	TypeOfPaintingCommonLife        = "common_life"
	TypeOfPaintingWar               = "war"
	TypeOfPaintingSceneWithGods     = "scene_with_gods"
	TypeOfPaintingPersonal          = "personal"
	TypeOfPaintingNonCourtGathering = "non_court_gathering"
)

func (p PaintingType) String() string {
	switch p {
	case RaagMala:
		return TypeOfPaintingRaagMala
	case CourtScene:
		return TypeOfPaintingCourtScene
	case CommonLife:
		return TypeOfPaintingCommonLife
	case War:
		return TypeOfPaintingWar
	case SceneWithGods:
		return TypeOfPaintingSceneWithGods
	case Personal:
		return TypeOfPaintingPersonal
	case NonCourtGathering:
		return TypeOfPaintingNonCourtGathering
	default:
		return "unknown"
	}
}

func GetPaintingTypeFromString(value string) (PaintingType, error) {
	switch value {
	case TypeOfPaintingRaagMala:
		return RaagMala, nil
	case TypeOfPaintingCourtScene:
		return CourtScene, nil
	case TypeOfPaintingCommonLife:
		return CommonLife, nil
	case TypeOfPaintingWar:
		return War, nil
	case TypeOfPaintingSceneWithGods:
		return SceneWithGods, nil
	case TypeOfPaintingPersonal:
		return Personal, nil
	case TypeOfPaintingNonCourtGathering:
		return NonCourtGathering, nil
	default:
		return -1, fmt.Errorf("invalid painting type value: %s", value)
	}
}
