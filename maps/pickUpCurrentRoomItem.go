package maps

import "escape-room-challenge/mapObjects"

func PickUpCurrentRoomItem(thisRoom *mapObjects.Room, systemMessage string, myItems []string) (string, []string) {
	switch thisRoom.ItemType {
	case 0:
	case 1:
		systemMessage = "열쇠를 획득했다."
		thisRoom.ItemType = 0
		myItems = append(myItems, "열쇠")
	case 2:
		systemMessage = "망치를 획득했다."
		thisRoom.ItemType = 0
		myItems = append(myItems, "망치")
	default:
	}
	return systemMessage, myItems
}
