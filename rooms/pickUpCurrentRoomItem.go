package rooms

import "escape-room-challenge/types"

func (room *Room) PickUpCurrentRoomItem(systemMessage string, myItems []string) (string, []string) {
	switch room.ItemType {
	case types.NoItem:
	case types.Key:
		systemMessage = "열쇠를 획득했다."
		room.ItemType = 0
		myItems = append(myItems, "열쇠")
	case types.Hammer:
		systemMessage = "망치를 획득했다."
		room.ItemType = 0
		myItems = append(myItems, "망치")
	default:
	}
	return systemMessage, myItems
}
