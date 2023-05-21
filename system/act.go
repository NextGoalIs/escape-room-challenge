package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/unit"
	"strings"
)

func Act(input string, char *unit.Character, systemMessage *string, isLookAtRoom *bool, stage *maps.MapStruct) {

	connectingRooms := stage.GetConnectingRooms()

	commandSlice := strings.Split(input, " ")
	firstCommand := ""
	secondCommand := ""
	thirdCommand := ""

	switch len(commandSlice) {
	case 1:
		firstCommand = commandSlice[0]
	case 2:
		firstCommand = commandSlice[0]
		secondCommand = commandSlice[1]
	case 3:
		firstCommand = commandSlice[0]
		secondCommand = commandSlice[1]
		thirdCommand = commandSlice[2]

	default:
	}

	switch secondCommand {
	case "사용":
		if maps.UseItem(firstCommand, &char.Items, connectingRooms, thirdCommand) {
			*systemMessage = "아이템을 사용했습니다."
			*isLookAtRoom = false
			return
		}
	case "열기":
		if maps.OpenDoor(firstCommand, connectingRooms, &char.Items) {
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
		if stage.GetThisLocation().PickUpItem(systemMessage, char) {
			*isLookAtRoom = false
			return
		}
		*systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
		return
	default:
		switch firstCommand {
		case "상태":
			char.ShowStatus()
			*systemMessage = ""
			return
		case "동", "서", "남", "북":
			if maps.Move(firstCommand, stage) {
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
