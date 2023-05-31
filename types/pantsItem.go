package types

type PantsItemTypes int

const (
	NoPantsItem PantsItemTypes = iota
	LeatherBottoms
)

var PantsDefenceMap = map[PantsItemTypes]int{
	NoPantsItem:    0,
	LeatherBottoms: 4,
}

var PantsNameMap = map[PantsItemTypes]string{
	NoPantsItem:    "",
	LeatherBottoms: "가죽바지",
}
