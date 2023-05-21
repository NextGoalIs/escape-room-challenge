package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"strings"
)

func UseItem(selectedCommand string, ableCommands []string, myItems *[]string, connectingRooms [4]*rooms.Room, thirdCommand string) bool {

	for _, item := range *myItems {
		if !strings.Contains(selectedCommand, item) {
			continue
		}

		for _, room := range connectingRooms {
			isUsed := useHammer(item, room, myItems, thirdCommand)
			if isUsed {
				return true
			}
		}

		for _, room := range connectingRooms {
			isUsed := useKey(item, room, myItems, thirdCommand)
			if isUsed {
				return true
			}
		}

	}

	return false
}

func useHammer(item string, room *rooms.Room, myItems *[]string, thirdCommand string) bool {
	if item != "망치" {
		return false
	}

	if room.DoorType != types.GlassType {
		return false
	}

	if thirdCommand != "유리문" {
		return false
	}

	room.SetEmpty()

	removeItem(myItems, item)
	return true
}

func useKey(item string, room *rooms.Room, myItems *[]string, thirdCommand string) bool {
	if item != "열쇠" {
		return false
	}

	if room.DoorType != types.LockedType {
		return false
	}

	if thirdCommand != "잠긴문" {
		return false
	}

	room.SetWoodDoor()

	removeItem(myItems, item)
	return true
}
