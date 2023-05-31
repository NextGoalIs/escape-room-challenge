package rooms

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
)

func (room *Room) PickUpItem(char *unit.Character, command string) types.UsingItemTypes {
	switch room.ItemType {
	case types.NoItem:
		return types.NoItem
	case types.Key:
		if command != "열쇠" {
			return types.NoItem
		}

		room.ItemType = 0
		char.Items = append(char.Items, "열쇠")

		return room.ItemType
	case types.Hammer:
		if command != "망치" {
			return types.NoItem
		}

		room.ItemType = 0
		char.Items = append(char.Items, "망치")

		return types.Hammer
	case types.DroppedWoodSword:
		if command != "목검" {
			return types.NoItem
		}

		room.ItemType = 0
		char.Items = append(char.Items, "목검")

		return types.DroppedWoodSword
	case types.Chest:
		if command != "상자" {
			return types.NoItem
		}

		room.ItemType = 0
		char.Items = append(char.Items, "상자")

		return types.Chest
	case types.Potion:
		if command != "회복약" {
			return types.NoItem
		}

		room.ItemType = 0
		char.Items = append(char.Items, "회복약")

		return types.Potion
	default:
		return types.NoItem
	}
}
