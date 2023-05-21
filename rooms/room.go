package rooms

import (
	"escape-room-challenge/types"
)

type Room struct {
	Icon string

	Unit     types.UnitTypes
	DoorType types.DoorTypes
	ItemType types.UsingItemTypes
	IsGoal   bool
}

func NewRoom(doorType types.DoorTypes, itemType types.UsingItemTypes, isGoal bool, unitType types.UnitTypes) Room {

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
	default:
		room.ItemType = itemType
	}

	switch unitType {
	case types.NoUnit:
	default:
		room.Unit = unitType
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
