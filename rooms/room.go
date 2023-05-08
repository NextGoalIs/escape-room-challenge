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
	case types.NoDoorType:
	case types.WoodType:
		room.Name = string(types.WoodName)
		room.DoorType = doorType
	case types.GlassType:
		room.Name = string(types.GlassName)
		room.DoorType = doorType
	case types.LockedType:
		room.Name = string(types.LockedName)
		room.DoorType = doorType
	default:
	}

	switch itemType {
	case types.NoItem:
	case types.Key:
		room.Name = string(types.KeyName)
		room.ItemType = itemType
	case types.Hammer:
		room.Name = string(types.HammerName)
		room.ItemType = itemType
	default:
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
