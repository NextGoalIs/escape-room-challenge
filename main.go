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

	defaultMap := maps.GetMap()
	nowX := 1
	nowY := 1
	var myItems []string

	defaultMap[nowX][nowY].SetMyCharacter()
	systemMessage := ""

	for {
		thisRoom := &defaultMap[nowX][nowY]
		var ableCommands []string

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}

		//무조건 먹게 하지 말고 줍게 하기...다음에...
		if thisRoom.ItemType != types.NoItem {
			thisRoom.PickUpItem(&systemMessage, &myItems)
		}

		// connectingRooms := [4]*rooms.Room{north, west, east, south}

		connectingRooms := maps.GetConnectingRooms(&defaultMap, nowX, nowY)

		var inputItem, inputCommand string

		system.AddMoveCommands(connectingRooms, &ableCommands)
		system.AddUseItemCommands(myItems, &ableCommands)
		system.AddOpenDoorCommands(connectingRooms, &ableCommands)

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		maps.Print(systemMessage, defaultMap, nowX, nowY, myItemsString, ableCommandsString, connectingRooms)
		fmt.Scanln(&inputItem, &inputCommand)

		var hasActed bool

		if inputCommand != "" && inputCommand == "사용" {
			hasActed = maps.UseItem(inputItem, ableCommandsString, nowX, nowY, &myItems, connectingRooms)

			if hasActed {
				systemMessage = "아이템을 사용했습니다."
				continue
			}
		}

		if inputCommand != "" && inputCommand == "열기" {
			hasActed = maps.OpenDoor(inputItem, ableCommandsString, &defaultMap, nowX, nowY, connectingRooms, &myItems)

			if hasActed {
				systemMessage = "문을 열었습니다."
				continue
			}
		}

		hasActed = maps.Move(inputItem, ableCommandsString, &defaultMap, &nowX, &nowY)
		systemMessage = ""

		if !hasActed {
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			continue
		}
	}

}
