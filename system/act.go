package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"strings"
)

func Act(char *unit.Character, message *Message, stage *maps.MapStruct) {

	input := Scan()

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
	case "전투":
		if !stage.GetIsViewedDetail() {
			message.SetCannotAct()
			return
		}

		if stage.GetThisLocation().Unit == types.NoUnit {
			message.SetCannotAct()
			return
		}

		// Battle(char, e)

	case "사용":
		if maps.UseItem(firstCommand, &char.Items, connectingRooms, thirdCommand) {
			message.SetUseItem()
			stage.SetIsViewedDetail(false)
			return
		}

		message.SetCannotAct()
	case "열기":
		if maps.OpenDoor(firstCommand, connectingRooms, &char.Items) {
			message.SetOpenDoor()
			stage.SetIsViewedDetail(false)
			return
		}

		message.SetCannotAct()
	case "보기":
		switch firstCommand {
		case "방":
			stage.SetIsViewedDetail(true)
			message.SetEmptyString()
			return
		default:
			message.SetCannotAct()
		}
		message.SetCannotAct()
		return
	case "줍기":
		if stage.GetThisLocation().PickUpItem(message, char) {
			stage.SetIsViewedDetail(false)
			return
		}
		message.SetCannotAct()
		return
	default:
		switch firstCommand {
		case "상태":
			char.ShowStatus()
			message.SetEmptyString()
			return
		case "동", "서", "남", "북":
			if maps.Move(firstCommand, stage) {
				stage.SetIsViewedDetail(false)
				message.SetEmptyString()
				return
			}

			stage.SetIsViewedDetail(false)
			message.SetCannotAct()
			return
		default:
			stage.SetIsViewedDetail(false)
			message.SetCannotAct()
			return
		}
	}
}
