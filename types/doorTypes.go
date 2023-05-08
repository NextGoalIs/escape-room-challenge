package types

type DoorTypes int
type DoorNames string

const (
	NoDoor DoorTypes = iota
	Wood
	Glass
	Locked
)

const (
	NoDoorName DoorNames = "🔳"
	WoodName   DoorNames = "🚪"
	GlassName  DoorNames = "🧊"
	LockedName DoorNames = "🔒"
	WallName   DoorNames = "🚧"
)
