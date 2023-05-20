package rooms

import (
	"escape-room-challenge/character"
	"escape-room-challenge/types"
)

func (room *Room) PickUpItem(systemMessage *string, char *character.Character) {
	switch room.ItemType {
	case types.NoItem:
	case types.Key:
		*systemMessage = "열쇠를 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "열쇠")
		// *myItems = append(*myItems, "열쇠")
	case types.Hammer:
		*systemMessage = "망치를 획득했다."
		room.ItemType = 0
		char.Items = append(char.Items, "망치")
		// *myItems = append(*myItems, "망치")
	default:
	}
}
