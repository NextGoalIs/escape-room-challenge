package rooms

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
)

func (room *Room) PickUpItem(systemMessage *string, char *unit.Character) bool {
	switch room.ItemType {
	case types.NoItem:
		return false
	case types.Key:
		*systemMessage = "열쇠를 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "열쇠")
		return true
	case types.Hammer:
		*systemMessage = "망치를 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "망치")
		return true
	case types.DroppedWoodSword:
		*systemMessage = "목검을 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "목검")
		return true
	case types.Chest:
		*systemMessage = "상자를 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "상자")
		return true
	case types.Potion:
		*systemMessage = "회복약을 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "회복약")
		return true
	default:
		return false
	}
}
