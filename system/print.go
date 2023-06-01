package system

import (
	"escape-room-challenge/maps"
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

func Print(stage maps.MapStruct, char unit.Character) {

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
		GetMessageInstance().SetItemsAreHere(stage.GetThisLocation().ItemType)

		enemy := ""
		if stage.GetThisLocation().Enemy.Name != "" && stage.GetThisLocation().Enemy.Health > 0 {
			enemy = types.EnemyIcon
			GetMessageInstance().SetEnemyIsHere(stage.GetThisLocation().Enemy.Name)
		}

		if firstThing == "" && enemy == "" {
			GetMessageInstance().SetRoomIsEmpty()
		}

		utils.ClearConsoleWindows()
		GetMessageInstance().Flush()
		println(utils.GetStringCenter(string(firstThing)+" "+north.Icon+" ", PADDING-len(west.Icon)))
		println(utils.GetStringCenter(west.Icon+" "+stage.GetNowMap()[stage.GetX()][stage.GetY()].Icon+" "+east.Icon, PADDING))
		println(utils.GetStringCenter(enemy+south.Icon, PADDING-len(west.Icon)))
		println()
		fmt.Println("소지품 : ", strings.Join(char.Items, ", "))
		fmt.Println("행동 : " + strings.Join(ableCommands, ", "))
		print(">>>  ")
	default:
		if stage.GetThisLocation().ItemType != types.NoItem || stage.GetThisLocation().Enemy.Name != "" {
			GetMessageInstance().SetSomethingInRoom()
		}

		utils.ClearConsoleWindows()
		GetMessageInstance().Flush()
		println(utils.GetStringCenter(north.Icon, PADDING-len(west.Icon)))
		println(utils.GetStringCenter(west.Icon+" "+stage.GetNowMap()[stage.GetX()][stage.GetY()].Icon+" "+east.Icon, PADDING))
		println(utils.GetStringCenter(south.Icon, PADDING-len(west.Icon)))
		println()
		fmt.Println("소지품 : ", strings.Join(char.Items, ", "))
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
