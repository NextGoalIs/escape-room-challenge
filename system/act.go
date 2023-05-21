package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/rooms"
	"escape-room-challenge/unit"
)

func Act(firstCommand string, secondCommand string, thirdCommand string, ableCommands []string, char *unit.Character, connectingRooms [4]*rooms.Room, systemMessage *string, isLookAtRoom *bool, stage1 *maps.MapStruct) {
	switch secondCommand {
	case "사용":
		if maps.UseItem(firstCommand, ableCommands, &char.Items, connectingRooms, thirdCommand) {
			*systemMessage = "아이템을 사용했습니다."
			*isLookAtRoom = false
			return
		}
	case "열기":
		if maps.OpenDoor(firstCommand, ableCommands, connectingRooms, &char.Items) {
			*systemMessage = "문을 열었습니다."
			*isLookAtRoom = false
			return
		}
	case "보기":
		switch firstCommand {
		case "방":
			*isLookAtRoom = true
			*systemMessage = ""
			return
		default:

		}
		*systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
		return
	case "줍기":
		if stage1.GetThisLocation().PickUpItem(systemMessage, char) {
			*isLookAtRoom = false
			return
		}
		*systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
		return
	default:
		switch firstCommand[0] {
		case "동"[0], "서"[0], "남"[0], "북"[0]:
			if maps.Move(firstCommand, ableCommands, stage1) {
				*isLookAtRoom = false
				*systemMessage = ""
				return
			}

			*isLookAtRoom = false
			*systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			return
		default:
			*isLookAtRoom = false
			*systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			return
		}
	}
}
