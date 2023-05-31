package rooms

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
)

type Room struct {
	Icon string

	Enemy    unit.Enemy
	DoorType types.DoorTypes
	ItemType types.UsingItemTypes
	IsGoal   bool
	IsWall   bool
}

func NewRoom(doorType types.DoorTypes, itemType types.UsingItemTypes, isGoal bool, enemyType types.UnitTypes) Room {

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

	switch enemyType {
	case types.Deer:
		room.Enemy = unit.NewDear()
	case types.Rabbit:
		room.Enemy = unit.NewRabbit()
	case types.Squirrel:
		room.Enemy = unit.NewSquirrel()
	case types.NoUnit:
	default:
	}

	if isGoal {
		room.IsGoal = isGoal
	}

	return room
}
