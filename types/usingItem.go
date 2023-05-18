package types

type UsingItemTypes int
type UsingItemIcons string

const (
	NoItem UsingItemTypes = iota
	Key
	Hammer
)

const (
	NoItemIcon UsingItemIcons = ""
	KeyIcon    UsingItemIcons = "🔑"
	HammerIcon UsingItemIcons = "🔨"
)
