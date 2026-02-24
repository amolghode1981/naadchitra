package types

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

type PaintingTypeEnum struct {
	Name        string
	IntValue    int
	StringValue PaintingTypeString
}

var PaintingTypes = []PaintingTypeEnum{
	{Name: "RaagMala", IntValue: RaagMala, StringValue: TypeOfPaintingRaagMala},
	{Name: "CourtScene", IntValue: CourtScene, StringValue: TypeOfPaintingCourtScene},
	{Name: "CommonLife", IntValue: CommonLife, StringValue: TypeOfPaintingCommonLife},
	{Name: "War", IntValue: War, StringValue: TypeOfPaintingWar},
	{Name: "SceneWithGods", IntValue: SceneWithGods, StringValue: TypeOfPaintingSceneWithGods},
	{Name: "Personal", IntValue: Personal, StringValue: TypeOfPaintingPersonal},
	{Name: "NonCourtGathering", IntValue: NonCourtGathering, StringValue: TypeOfPaintingNonCourtGathering},
}

var PaintingTypeByInt = make(map[int]PaintingTypeEnum)
var PaintingTypeByString = make(map[PaintingTypeString]PaintingTypeEnum)

func init() {
	for _, pt := range PaintingTypes {
		PaintingTypeByInt[pt.IntValue] = pt
		PaintingTypeByString[pt.StringValue] = pt
	}
}
