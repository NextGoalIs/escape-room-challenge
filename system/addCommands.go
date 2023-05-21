package system

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
)

func AddCommands(connectingRooms [4]*rooms.Room, ableCommands *[]string, char unit.Character) {
	AddMoveCommands(connectingRooms, ableCommands)
	AddUseItemCommands(char.Items, ableCommands)
	AddOpenDoorCommands(connectingRooms, ableCommands)
	AddDefaultCommands(ableCommands)
}

func AddMoveCommands(connectingRooms [4]*rooms.Room, ableCommands *[]string) {
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

func AddUseItemCommands(myItems []string, ableCommands *[]string) {
	for _, v := range myItems {
		*ableCommands = append(*ableCommands, v+" 사용")
	}
}

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

func AddDefaultCommands(ableCommands *[]string) {
	*ableCommands = append(*ableCommands, "<목표> 보기")
	*ableCommands = append(*ableCommands, "<목표> 줍기")
}
