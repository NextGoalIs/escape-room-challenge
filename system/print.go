package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func Print(systemMessage *string, stage maps.MapStruct, char unit.Character) {

	var ableCommands []string
	AddCommands(stage, &ableCommands, char)

	const PADDING int = 60

	connectingRooms := stage.GetConnectingRooms()
	north := connectingRooms[3]
	west := connectingRooms[1]
	east := connectingRooms[0]
	south := connectingRooms[2]

	switch stage.GetIsViewedDetail() {
	case true:
		firstThing := setIcon(stage.GetThisLocation().ItemType)

		enemy := ""
		if stage.GetThisLocation().Unit != types.NoUnit {
			enemy = types.EnemyIcon
		}

		*systemMessage = setSystemItemMessage(stage.GetThisLocation().ItemType)

		if firstThing == "" && enemy == "" {
			*systemMessage = "이 방엔 아무것도 없는 듯 하다..."
		}

		utils.ClearConsoleWindows()
		fmt.Println(*systemMessage)
		println(utils.GetStringCenter(string(firstThing)+" "+north.Icon+" ", PADDING-len(west.Icon)))
		println(utils.GetStringCenter(west.Icon+" "+stage.GetNowMap()[stage.GetX()][stage.GetY()].Icon+" "+east.Icon, PADDING))
		println(utils.GetStringCenter(enemy+south.Icon, PADDING-len(west.Icon)))
		println()
		fmt.Println("행동 : " + strings.Join(ableCommands, ", "))
		print(">>>  ")
	default:
		if stage.GetThisLocation().ItemType != types.NoItem || stage.GetThisLocation().Unit != types.NoUnit {
			*systemMessage = "이 방에 무언가 있는 것 같다."
		}

		utils.ClearConsoleWindows()
		fmt.Println(*systemMessage)
		println(utils.GetStringCenter(north.Icon, PADDING-len(west.Icon)))
		println(utils.GetStringCenter(west.Icon+" "+stage.GetNowMap()[stage.GetX()][stage.GetY()].Icon+" "+east.Icon, PADDING))
		println(utils.GetStringCenter(south.Icon, PADDING-len(west.Icon)))
		println()
		fmt.Println("행동 : " + strings.Join(ableCommands, ", "))
		print(">>>  ")
	}
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
	case types.DroppedWoodSword:
		return types.DroppedWoodSwordIcon
	default:
		return types.NoItemIcon
	}
}

func setSystemItemMessage(itemType types.UsingItemTypes) string {
	switch itemType {
	case types.Chest:
		return "방에 상자가 있다."
	case types.Hammer:
		return "방에 망치가 있다."
	case types.Key:
		return "방에 열쇠가 있다."
	case types.Potion:
		return "방에 회복약이 있다."
	case types.DroppedWoodSword:
		return "방에 목검이 있다."
	default:
		return ""
	}
}
