package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"strings"
)

func UseItem(selectedCommand string, ableCommandsString string, nowX int, nowY int, myItems *[]string, connectingRooms [4]*rooms.Room) bool {
	if !strings.Contains(ableCommandsString, selectedCommand) {
		return false
	}

	for _, item := range *myItems {
		if !strings.Contains(selectedCommand, item) {
			continue
		}

		for _, room := range connectingRooms {
			isUsed := useHammer(item, room, myItems)
			if isUsed {
				return true
			}
		}

		for _, room := range connectingRooms {
			isUsed := useKey(item, room, myItems)
			if isUsed {
				return true
			}
		}

	}

	return false
}

func useHammer(item string, room *rooms.Room, myItems *[]string) bool {
	if item != "망치" {
		return false
	}

	if room.DoorType != types.GlassType {
		return false
	}

	room.SetEmptyRoom()

	removeItem(myItems, item)
	return true
}

func useKey(item string, room *rooms.Room, myItems *[]string) bool {
	if item != "열쇠" {
		return false
	}

	if room.DoorType != types.LockedType {
		return false
	}

	room.SetWoodDoor()

	removeItem(myItems, item)
	return true
}
