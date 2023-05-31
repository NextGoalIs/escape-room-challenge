package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"strings"
)

func Act(char *unit.Character, stage *maps.MapStruct) {

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
			GetMessageInstance().SetCannotAct()
			return
		}

		if stage.GetThisLocation().Enemy.Name == "" {
			GetMessageInstance().SetCannotAct()
			return
		}

		Battle(char, &stage.GetThisLocation().Enemy)

	case "사용":
		if maps.UseItem(firstCommand, &char.Items, connectingRooms, thirdCommand) {
			GetMessageInstance().SetUseItem()
			stage.SetIsViewedDetail(false)
			return
		}

		GetMessageInstance().SetCannotAct()
	case "열기":
		if maps.OpenDoor(firstCommand, connectingRooms, &char.Items) {
			GetMessageInstance().SetOpenDoor()
			stage.SetIsViewedDetail(false)
			return
		}

		GetMessageInstance().SetCannotAct()
	case "보기":
		switch firstCommand {
		case "방":
			stage.SetIsViewedDetail(true)
			GetMessageInstance().SetEmptyString()
			return
		default:
			GetMessageInstance().SetCannotAct()
		}
		GetMessageInstance().SetCannotAct()
		return
	case "줍기":
		if stage.GetThisLocation().ItemType == types.NoItem {
			GetMessageInstance().SetCannotAct()
			return
		}

		GetMessageInstance().SetPickUpItem(stage.GetThisLocation().PickUpItem(char, firstCommand))
		return
	default:
		switch firstCommand {
		case "상태":
			char.ShowStatus()
			GetMessageInstance().SetEmptyString()
			return
		case "동", "서", "남", "북":
			if maps.Move(firstCommand, stage) {
				stage.SetIsViewedDetail(false)
				GetMessageInstance().SetEmptyString()
				return
			}

			stage.SetIsViewedDetail(false)
			GetMessageInstance().SetCannotAct()
			return
		default:
			stage.SetIsViewedDetail(false)
			GetMessageInstance().SetCannotAct()
			return
		}
	}
}
