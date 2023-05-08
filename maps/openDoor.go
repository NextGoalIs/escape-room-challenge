package maps

import (
	"escape-room-challenge/rooms"
	"strings"
)

func OpenDoor(inputItem string, ableCommandsString string, defaultMap *[6][8]rooms.Room, nowX int, nowY int, directionRoom [4]*rooms.Room, myItems *[]string) bool {
	if !strings.Contains(ableCommandsString, inputItem) {
		return false
	}

	switch inputItem {
	case "나무문":
		for _, room := range directionRoom {
			if room.DoorType != 1 {
				continue
			}

			room.SetEmptyRoom()
			break
		}
	case "유리문":
		if !strings.Contains(ableCommandsString, "망치 사용") {
			return false
		}

		for _, room := range directionRoom {
			if room.DoorType != 2 {
				continue
			}

			room.SetEmptyRoom()
			removeItem(myItems, "망치")
			break
		}
	case "잠긴문":
		if !strings.Contains(ableCommandsString, "열쇠 사용") {
			return false
		}

		for _, room := range directionRoom {
			if room.DoorType != 3 {
				continue
			}

			setWoodDoor(room)
			break
		}
	default:
		return false
	}

	return true
}
