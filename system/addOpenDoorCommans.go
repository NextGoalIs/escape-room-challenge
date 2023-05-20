package system

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func AddOpenDoorCommands(connectingRooms [4]*rooms.Room, ableCommands *[]string) {
	for _, v := range connectingRooms {
		if v.DoorType == types.NoDoorType {
			continue
		}

		if v.DoorType == types.WoodType {
			*ableCommands = append(*ableCommands, "나무문 열기")
			continue
		}

		if v.DoorType == types.GlassType {
			*ableCommands = append(*ableCommands, "유리문 열기")
			continue
		}

		if v.DoorType == types.LockedType {
			*ableCommands = append(*ableCommands, "잠긴문 열기")
			continue
		}
	}
}
