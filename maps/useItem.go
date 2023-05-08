package maps

import (
	"escape-room-challenge/mapObjects"
	"strings"
)

func UseItem(selectedCommand string, ableCommandsString string, nowX int, nowY int, myItems *[]string, directionRoom [4]*mapObjects.Room) bool {
	if !strings.Contains(ableCommandsString, selectedCommand) {
		return false
	}

	for _, item := range *myItems {
		if !strings.Contains(selectedCommand, item) {
			continue
		}

		for _, room := range directionRoom {
			isUsed := useHammer(item, room, myItems)
			if isUsed {
				return true
			}
		}

		for _, room := range directionRoom {
			isUsed := useKey(item, room, myItems)
			if isUsed {
				return true
			}
		}

	}

	return false
}

func removeItem(myItems *[]string, item string) {
	if len(*myItems) == 1 && (*myItems)[0] == item {
		*myItems = []string{}
	} else {
		for i, myItem := range *myItems {
			if myItem == item {
				*myItems = append((*myItems)[:i], (*myItems)[i+1])
				break
			}
		}
	}
}

func useHammer(item string, room *mapObjects.Room, myItems *[]string) bool {
	if item != "망치" {
		return false
	}

	if room.DoorType != 2 {
		return false
	}

	room.DoorType = 0
	room.Name = "🔳"

	removeItem(myItems, item)
	return true
}

func useKey(item string, room *mapObjects.Room, myItems *[]string) bool {
	if item != "열쇠" {
		return false
	}

	if room.DoorType != 3 {
		return false
	}

	room.DoorType = 0
	room.Name = "🔳"

	removeItem(myItems, item)
	return true
}
