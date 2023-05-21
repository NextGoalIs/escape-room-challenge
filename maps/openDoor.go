package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
)

func OpenDoor(inputItem string, ableCommands []string, connectingRooms [4]*rooms.Room, myItems *[]string) bool {

	switch inputItem {
	case "나무문":
		for _, room := range connectingRooms {
			if room.DoorType != types.WoodType {
				continue
			}

			room.SetEmpty()
			break
		}
	case "유리문":
		if !utils.Contains(ableCommands, "망치 사용") {
			return false
		}

		for _, room := range connectingRooms {
			if room.DoorType != types.GlassType {
				continue
			}

			room.SetEmpty()
			removeItem(myItems, "망치")
			break
		}
	case "잠긴문":
		if !utils.Contains(ableCommands, "열쇠 사용") {
			return false
		}

		for _, room := range connectingRooms {
			if room.DoorType != types.LockedType {
				continue
			}

			room.SetWoodDoor()
			removeItem(myItems, "열쇠")
			break
		}
	default:
		return false
	}

	return true
}
