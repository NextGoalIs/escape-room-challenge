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

func useHammer(item string, room *mapObjects.Room, myItems *[]string) bool {
	if item != "망치" {
		return false
	}

	if room.DoorType != 2 {
		return false
	}

	setEmptyRoom(room)

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

	setWoodDoor(room)

	removeItem(myItems, item)
	return true
}
