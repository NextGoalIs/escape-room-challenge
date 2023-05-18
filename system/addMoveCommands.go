package system

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func AddMoveCommands(connectingRooms [4]rooms.Room, ableCommands *[]string) {
	for i, v := range connectingRooms {
		if v.Icon != string(types.WallIcon) && v.DoorType == types.NoDoorType {
			switch i {
			case int(types.East):
				*ableCommands = append(*ableCommands, "동")
			case int(types.West):
				*ableCommands = append(*ableCommands, "서")
			case int(types.South):
				*ableCommands = append(*ableCommands, "남")
			case int(types.North):
				*ableCommands = append(*ableCommands, "북")
			default:
			}
		}
	}
}
