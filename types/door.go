package types

type DoorTypes int
type DoorIcons string

const (
	NoDoorType DoorTypes = iota
	WoodType
	GlassType
	LockedType
)

const (
	NoDoorIcon      DoorIcons = "🔳"
	WoodIcon        DoorIcons = "🚪"
	GlassIcon       DoorIcons = "🧊"
	LockedIcon      DoorIcons = "🔒"
	WallIcon        DoorIcons = "🚧"
	MyCharacterIcon DoorIcons = "🏃"
)
