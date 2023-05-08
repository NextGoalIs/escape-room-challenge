package main

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/maps"
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

	PADDING := 60

	defaultMap[nowX][nowY].Name = "🏃"

	for {

		systemMessage := ""

		thisRoom := &defaultMap[nowX][nowY]

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}
	CommandSwitch:
		systemMessage, myItems = maps.PickUpCurrentRoomItem(thisRoom, systemMessage, myItems)

		east, ableCommands, west, south, north := maps.MakeConsoleMap(&defaultMap, nowX, nowY)
		directionRoom := [4]*mapObjects.Room{north, west, east, south}

		var inputItem, inputCommand string

		for _, v := range myItems {
			ableCommands = append(ableCommands, v+" 사용")
		}

		for _, v := range directionRoom {
			if v.DoorType == 0 {
				continue
			}

			if v.DoorType == types.Wood {
				ableCommands = append(ableCommands, "나무문 열기")
				continue
			}

			if v.DoorType == types.Glass {
				ableCommands = append(ableCommands, "유리문 열기")
				continue
			}

			if v.DoorType == types.Locked {
				ableCommands = append(ableCommands, "잠긴문 열기")
				continue
			}
		}

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		maps.PrintDisplay(systemMessage, PADDING, defaultMap, nowX, nowY, myItemsString, ableCommandsString, directionRoom)
		fmt.Scanln(&inputItem, &inputCommand)

		var hasActed bool

		if inputCommand != "" && inputCommand == "사용" {
			hasActed = maps.UseItem(inputItem, ableCommandsString, nowX, nowY, &myItems, directionRoom)

			if hasActed {
				systemMessage = "아이템을 사용했습니다."
				goto CommandSwitch
			}
		}

		if inputCommand != "" && inputCommand == "열기" {
			hasActed = maps.OpenDoor(inputItem, ableCommandsString, &defaultMap, nowX, nowY, directionRoom, &myItems)

			if hasActed {
				systemMessage = "문을 열었습니다."
				goto CommandSwitch
			}
		}

		hasActed = maps.Move(inputItem, ableCommandsString, &defaultMap, &nowX, &nowY)

		if !hasActed {
			systemMessage = "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다."
			goto CommandSwitch
		}
	}

}
