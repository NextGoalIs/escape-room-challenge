package types

type DoorTypes int
type DoorNames string

const (
	NoDoorType DoorTypes = iota
	WoodType
	GlassType
	LockedType
)

const (
	NoDoorName      DoorNames = "🔳"
	WoodName        DoorNames = "🚪"
	GlassName       DoorNames = "🧊"
	LockedName      DoorNames = "🔒"
	WallName        DoorNames = "🚧"
	MyCharacterName DoorNames = "🏃"
)
