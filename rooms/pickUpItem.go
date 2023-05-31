package rooms

import (
	"escape-room-challenge/system"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
)

func (room *Room) PickUpItem(message *system.Message, char *unit.Character) bool {
	switch room.ItemType {
	case types.NoItem:
		return false
	case types.Key:
		message.SetPickUpItem(room.ItemType)
		room.ItemType = 0
		char.Items = append(char.Items, "열쇠")
		return true
	case types.Hammer:
		message.SetPickUpItem(room.ItemType)
		room.ItemType = 0
		char.Items = append(char.Items, "망치")
		return true
	case types.DroppedWoodSword:
		message.SetPickUpItem(room.ItemType)
		room.ItemType = 0
		char.Items = append(char.Items, "목검")
		return true
	case types.Chest:
		message.SetPickUpItem(room.ItemType)
		room.ItemType = 0
		char.Items = append(char.Items, "상자")
		return true
	case types.Potion:
		message.SetPickUpItem(room.ItemType)
		room.ItemType = 0
		char.Items = append(char.Items, "회복약")
		return true
	default:
		return false
	}
}
