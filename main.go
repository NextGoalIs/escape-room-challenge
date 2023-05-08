package main

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/maps"
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

		_, ok := interface{}(defaultMap[nowX][nowY]).(mapObjects.Room)
		if !ok {
			fmt.Println("Error")
			break
		}

		thisRoom := &defaultMap[nowX][nowY]

		if thisRoom.IsGoal {
			utils.PrintWIN()
			break
		}
	CommandSwitch:
		systemMessage, myItems = maps.PickUpCurrentRoomItem(thisRoom, systemMessage, myItems)

		east, ableCommands, west, south, north := maps.MakeConsoleMap(&defaultMap, nowX, nowY)

		var inputItem, inputCommand string

		for _, v := range myItems {
			ableCommands = append(ableCommands, v+" 사용")
		}

		ableCommandsString := strings.Join(ableCommands, ", ")
		myItemsString := strings.Join(myItems, ", ")

		//시스템
		maps.PrintDisplay(systemMessage, north, PADDING, west, defaultMap, nowX, nowY, east, south, myItemsString, ableCommandsString)
		fmt.Scanln(&inputItem, &inputCommand)

		var hasActed bool

		if inputCommand != "" && inputCommand == "사용" {
			hasActed = maps.UseItem(inputItem, ableCommandsString, nowX, nowY, &myItems, east, west, south, north)

			if hasActed {
				systemMessage = "아이템을 사용했습니다."
				goto CommandSwitch
			}
		}

		hasActed = maps.Move(inputItem, ableCommandsString, &defaultMap, &nowX, &nowY)

		if !hasActed {
			systemMessage = "할 수 없는 행동입니다."
			goto CommandSwitch
		}

	}

}
