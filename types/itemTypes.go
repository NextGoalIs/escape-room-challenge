package types

type ItemTypes int
type ItemNames string

const (
	NoItem ItemTypes = iota
	Key
	Hammer
)

const (
	NoItemName ItemNames = ""
	KeyName    ItemNames = "🔑"
	HammerName ItemNames = "🔨"
)
