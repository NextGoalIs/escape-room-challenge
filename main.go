package main

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/system"
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func main() {

	stage1 := maps.GetStage1()
	var myItems []string

	stage1.GetThisLocation().SetMyCharacter()
	systemMessage := ""

	for {
		var ableCommands []string
		var inputItem, inputCommand string

		if stage1.GetThisLocation().IsGoal {
			utils.PrintWIN()
			break
		}

		//무조건 먹게 하지 말고 줍게 하기...다음에...
		if stage1.GetThisLocation().ItemType != types.NoItem {
			stage1.GetThisLocation().PickUpItem(&systemMessage, &myItems)
		}

		connectingRooms := stage1.GetConnectingRooms()

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(myItems, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		maps.Print(systemMessage, stage1, myItemsString, ableCommandsString, connectingRooms)
		fmt.Scanln(&inputItem, &inputCommand)

		switch inputCommand {
		case "사용":
			if maps.UseItem(inputItem, ableCommandsString, &myItems, connectingRooms) {
				systemMessage = "아이템을 사용했습니다."
				continue
			}
		case "열기":
			if maps.OpenDoor(inputItem, ableCommandsString, stage1, connectingRooms, &myItems) {
				systemMessage = "문을 열었습니다."
				continue
			}
		default:
			switch inputItem[0] {
			case "동"[0], "서"[0], "남"[0], "북"[0]:
				if maps.Move(inputItem, ableCommandsString, &stage1) {
					systemMessage = ""
					continue
				}

				fallthrough
			default:
				systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
				continue
			}
		}
	}

}
