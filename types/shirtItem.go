package types

type ShirtItemTypes int

const (
	NoShirtItem ShirtItemTypes = iota
	LeatherJacket
)

var ShirtDefenceMap = map[ShirtItemTypes]int{
	NoShirtItem:   0,
	LeatherJacket: 6,
}

var ShirtNameMap = map[ShirtItemTypes]string{
	NoShirtItem:   "",
	LeatherJacket: "가죽옷",
}
