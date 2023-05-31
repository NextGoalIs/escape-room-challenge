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

var WeaponNameMap = map[WeaponTypes]string{
	NoWeaponItem: "",
	WoodSword:    "목검",
	IronSword:    "철검",
}
