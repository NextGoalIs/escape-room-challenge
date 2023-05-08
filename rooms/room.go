package rooms

import "escape-room-challenge/types"

type Room struct {
	Name string

	DoorType types.DoorTypes
	ItemType types.ItemTypes
	IsGoal   bool
}

func NewRoom(doorType types.DoorTypes, itemType types.ItemTypes, isGoal bool) Room {

	room := Room{Name: "🔳"}

	switch doorType {
	case 0:
	case 1:
		room.Name = "🚪"
		room.DoorType = doorType
	case 2:
		room.Name = "🧊"
		room.DoorType = doorType
	case 3:
		room.Name = "🔒"
		room.DoorType = doorType
	default:
	}

	switch itemType {
	case 0:
	case 1:
		room.Name = "🔑"
		room.ItemType = itemType
	case 2:
		room.Name = "🔨"
		room.ItemType = itemType
	default:
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
