package maps

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/utils"
	"fmt"
)

func PrintDisplay(systemMessage string, north *mapObjects.Room, PADDING int, west *mapObjects.Room, defaultMap [6][8]mapObjects.Room, nowX int, nowY int, east *mapObjects.Room, south *mapObjects.Room, myItemsString string, ableCommandsString string) {
	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(north.GetName(), PADDING-len(west.GetName())))
	println(utils.GetStringCenter(west.GetName()+" "+defaultMap[nowX][nowY].GetName()+" "+east.GetName(), PADDING))
	println(utils.GetStringCenter(south.GetName(), PADDING-len(west.GetName())))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("할 수 있는 행동 : " + ableCommandsString)

	print(">>>  ")
}
