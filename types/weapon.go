package types

type WeaponTypes int

const (
	NoWeaponItem = iota
	WoodSword
	IronSword
	WoodShield
)

var WeaponDamageMap = map[WeaponTypes]int{
	NoWeaponItem: 0,
	WoodSword:    5,
	IronSword:    10,
}
