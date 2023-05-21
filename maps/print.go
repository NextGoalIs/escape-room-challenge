package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

const PADDING int = 60

func Print(systemMessage string, stage mapStruct, myItemsString string, ableCommands []string, connectingRooms [4]*rooms.Room) {
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
	fmt.Println("행동 : " + strings.Join(ableCommands, ", "))

	print(">>>  ")
}

func LookAtRoomPrint(systemMessage string, stage mapStruct, myItemsString string, ableCommands []string, connectingRooms [4]*rooms.Room) {
	north := connectingRooms[3]
	west := connectingRooms[1]
	east := connectingRooms[0]
	south := connectingRooms[2]

	firstThing := setIcon(stage.GetThisLocation().ItemType)

	// secondThing := setIcon(stage.)

	systemMessage = setSystemMessage(stage.GetThisLocation().ItemType)

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
	fmt.Println("행동 : " + strings.Join(ableCommands, ", "))

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

func setSystemMessage(itemType types.UsingItemTypes) string {
	switch itemType {
	case types.Chest:
		return "방에 상자가 있다."
	case types.Hammer:
		return "방에 망치가 있다."
	case types.Key:
		return "방에 열쇠가 있다."
	case types.Potion:
		return "방에 포션이 있다."
	case types.WoodSword:
		return "방에 목검이 있다."
	default:
		return ""
	}
}
