package rooms

import "escape-room-challenge/types"

type Room struct {
	Icon string

	DoorType types.DoorTypes
	ItemType types.UsingItemTypes
	IsGoal   bool
}

func NewRoom(doorType types.DoorTypes, itemType types.UsingItemTypes, isGoal bool) Room {

	room := Room{Icon: "🔳"}

	switch doorType {
	case types.NoDoorType:
	case types.WoodType:
		room.Icon = string(types.WoodIcon)
		room.DoorType = doorType
	case types.GlassType:
		room.Icon = string(types.GlassIcon)
		room.DoorType = doorType
	case types.LockedType:
		room.Icon = string(types.LockedIcon)
		room.DoorType = doorType
	default:
	}

	switch itemType {
	case types.NoItem:
	case types.Key:
		room.Icon = string(types.KeyIcon)
		room.ItemType = itemType
	case types.Hammer:
		room.Icon = string(types.HammerIcon)
		room.ItemType = itemType
	default:
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
