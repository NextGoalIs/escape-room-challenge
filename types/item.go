package types

type ItemTypes int
type ItemIcons string

const (
	NoItem ItemTypes = iota
	Key
	Hammer
)

const (
	NoItemIcon ItemIcons = ""
	KeyIcon    ItemIcons = "🔑"
	HammerIcon ItemIcons = "🔨"
)
