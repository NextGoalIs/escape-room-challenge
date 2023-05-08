package maps

import (
	"escape-room-challenge/mapObjects"
	"escape-room-challenge/utils"
	"fmt"
)

func PrintDisplay(systemMessage string, PADDING int, defaultMap [6][8]mapObjects.Room, nowX int, nowY int, myItemsString string, ableCommandsString string, directionRoom [4]*mapObjects.Room) {
	north := directionRoom[0]
	west := directionRoom[1]
	east := directionRoom[2]
	south := directionRoom[3]

	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(north.GetName(), PADDING-len(west.GetName())))
	println(utils.GetStringCenter(west.GetName()+" "+defaultMap[nowX][nowY].GetName()+" "+east.GetName(), PADDING))
	println(utils.GetStringCenter(south.GetName(), PADDING-len(west.GetName())))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("행동 : " + ableCommandsString)

	print(">>>  ")
}
