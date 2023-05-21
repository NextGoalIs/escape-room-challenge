package types

type UsingItemTypes int
type UsingItemIcons string

const (
	NoItem UsingItemTypes = iota
	Key
	Hammer
	Chest
	Potion
	WoodSword
)

const (
	NoItemIcon    UsingItemIcons = ""
	KeyIcon       UsingItemIcons = "🔑"
	HammerIcon    UsingItemIcons = "🔨"
	ChestIcon     UsingItemIcons = "🎁"
	PotionIcon    UsingItemIcons = "🧪"
	WoodSwordIcon UsingItemIcons = "🔪"
)
