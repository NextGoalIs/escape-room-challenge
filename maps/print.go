package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/utils"
	"fmt"
)

const PADDING int = 60

func Print(systemMessage string, stage mapStruct, myItemsString string, ableCommandsString string, connectingRooms [4]*rooms.Room) {
	north := connectingRooms[3]
	west := connectingRooms[1]
	east := connectingRooms[0]
	south := connectingRooms[2]

	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(north.Icon, PADDING-len(west.Icon)))
	println(utils.GetStringCenter(west.Icon+" "+stage.nowMap[stage.nowX][stage.nowY].Icon+" "+east.Icon, PADDING))
	println(utils.GetStringCenter(south.Icon, PADDING-len(west.Icon)))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("행동 : " + ableCommandsString)

	print(">>>  ")
}
