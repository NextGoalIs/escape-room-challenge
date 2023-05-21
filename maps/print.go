package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
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

func LookAtRoomPrint(systemMessage string, stage mapStruct, myItemsString string, ableCommandsString string, connectingRooms [4]*rooms.Room) {
	north := connectingRooms[3]
	west := connectingRooms[1]
	east := connectingRooms[0]
	south := connectingRooms[2]

	firstThing := setIcon(stage.GetThisLocation().ItemType)
	// secondThing := setIcon(stage.)

	if firstThing == "" {
		systemMessage = "이 방엔 아무것도 없는 듯 하다..."
	}

	utils.ClearConsoleWindows()
	fmt.Println(systemMessage)
	println(utils.GetStringCenter(string(firstThing)+" "+north.Icon+" ", PADDING-len(west.Icon)))
	println(utils.GetStringCenter(west.Icon+" "+stage.nowMap[stage.nowX][stage.nowY].Icon+" "+east.Icon, PADDING))
	println(utils.GetStringCenter(" "+south.Icon+" ", PADDING-len(west.Icon)))
	println()
	fmt.Println("가지고 있는 물건 : " + myItemsString)
	fmt.Println("행동 : " + ableCommandsString)

	print(">>>  ")
}

func setIcon(itemType types.UsingItemTypes) types.UsingItemIcons {
	switch itemType {
	case types.Key:
		return types.KeyIcon
	case types.Hammer:
		return types.HammerIcon
	case types.Chest:
		return types.ChestIcon
	case types.Potion:
		return types.PotionIcon
	case types.WoodSword:
		return types.WoodSwordIcon
	default:
		return types.NoItemIcon
	}
}
