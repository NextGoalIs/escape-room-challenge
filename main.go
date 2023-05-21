package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func main() {

	char := unit.Character{}
	char.SetName()

	stage1 := maps.GetStage1()

	stage1.GetThisLocation().SetMyCharacter()
	systemMessage := ""
	isLookAtRoom := false

	for {
		var ableCommands []string
		var firstCommand, secondCommand string

		if stage1.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		connectingRooms := stage1.GetConnectingRooms()

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(char.Items, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(char.Items, ", ")

		switch isLookAtRoom {
		case true:
			maps.LookAtRoomPrint(systemMessage, stage1, myItemsString, ableCommandsString, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand)
		default:
			maps.Print(systemMessage, stage1, myItemsString, ableCommandsString, connectingRooms)
			fmt.Scanln(&firstCommand, &secondCommand)
		}

		switch secondCommand {
		case "사용":
			if maps.UseItem(firstCommand, ableCommandsString, &char.Items, connectingRooms) {
				systemMessage = "아이템을 사용했습니다."
				isLookAtRoom = false
				continue
			}
		case "열기":
			if maps.OpenDoor(firstCommand, ableCommandsString, connectingRooms, &char.Items) {
				systemMessage = "문을 열었습니다."
				isLookAtRoom = false
				continue
			}
		case "보기":
			if char.LookAt(firstCommand) {
				isLookAtRoom = true
				systemMessage = ""
				continue
			}
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			continue
		case "줍기":
			if stage1.GetThisLocation().PickUpItem(&systemMessage, &char) {
				isLookAtRoom = false
				continue
			}
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			continue
		default:
			switch firstCommand[0] {
			case "동"[0], "서"[0], "남"[0], "북"[0]:
				if maps.Move(firstCommand, ableCommandsString, &stage1) {
					isLookAtRoom = false
					systemMessage = ""
					continue
				}

				fallthrough
			default:
				isLookAtRoom = false
				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
				continue
			}
		}
	}

}
