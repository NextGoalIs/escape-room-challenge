package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
	"math/rand"
	"strings"
	"time"
)

func UseItem(selectedCommand string, myItems *[]string, connectingRooms [4]*rooms.Room, thirdCommand string) bool {

	for _, item := range *myItems {
		if !strings.Contains(selectedCommand, item) {
			continue
		}

		if item == "상자" {
			pushChestItemRandomly(myItems)
			RemoveItem(myItems, item)
			return true
		}

		for _, room := range connectingRooms {
			isUsed := useHammer(item, room, myItems, thirdCommand)
			if isUsed {
				return true
			}
		}

		for _, room := range connectingRooms {
			isUsed := useKey(item, room, myItems, thirdCommand)
			if isUsed {
				return true
			}
		}

	}

	return false
}

func useHammer(item string, room *rooms.Room, myItems *[]string, thirdCommand string) bool {
	if item != "망치" {
		return false
	}

	if room.DoorType != types.GlassType {
		return false
	}

	if thirdCommand != "유리문" {
		return false
	}

	room.SetEmpty()

	RemoveItem(myItems, item)
	return true
}

func useKey(item string, room *rooms.Room, myItems *[]string, thirdCommand string) bool {
	if item != "열쇠" {
		return false
	}

	if room.DoorType != types.LockedType {
		return false
	}

	if thirdCommand != "잠긴문" {
		return false
	}

	room.SetWoodDoor()

	RemoveItem(myItems, item)
	return true
}

func pushChestItemRandomly(myItems *[]string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	switch n := rand.Intn(100); {
	case n < 20:
		*myItems = append(*myItems, "목검")
	case n < 36:
		*myItems = append(*myItems, "철검")
	case n < 43:
		*myItems = append(*myItems, "가죽옷")
	case n < 51:
		*myItems = append(*myItems, "가죽바지")
	case n < 61:
		*myItems = append(*myItems, "가죽신발")
	case n < 76:
		*myItems = append(*myItems, "회복약")
	case n < 86:
		*myItems = append(*myItems, "회복약")
		*myItems = append(*myItems, "회복약")
	case n < 91:
		*myItems = append(*myItems, "회복약")
		*myItems = append(*myItems, "회복약")
		*myItems = append(*myItems, "회복약")
	default:
	}
}
