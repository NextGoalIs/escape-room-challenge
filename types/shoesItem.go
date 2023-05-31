package types

type ShoesItemTypes int

const (
	NoShoesItem ShoesItemTypes = iota
	LeatherShoes
)

var ShoesDefenceMap = map[ShoesItemTypes]int{
	NoShoesItem:  0,
	LeatherShoes: 3,
}

var ShoesNameMap = map[ShoesItemTypes]string{
	NoShoesItem:  "",
	LeatherShoes: "가죽신발",
}
