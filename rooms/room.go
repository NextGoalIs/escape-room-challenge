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
	case types.NoDoor:
	case types.Wood:
		room.Name = "🚪"
		room.DoorType = doorType
	case types.Glass:
		room.Name = "🧊"
		room.DoorType = doorType
	case types.Locked:
		room.Name = "🔒"
		room.DoorType = doorType
	default:
	}

	switch itemType {
	case types.NoItem:
	case types.Key:
		room.Name = "🔑"
		room.ItemType = itemType
	case types.Hammer:
		room.Name = "🔨"
		room.ItemType = itemType
	default:
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
