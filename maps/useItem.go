package maps

import (
	"escape-room-challenge/mapObjects"
	"strings"
)

func UseItem(selectedCommand string, ableCommandsString string, nowX int, nowY int, myItems *[]string, east *mapObjects.Room, west *mapObjects.Room, south *mapObjects.Room, north *mapObjects.Room) bool {
	if !strings.Contains(ableCommandsString, selectedCommand) {
		return false
	}

	for _, item := range *myItems {
		if !strings.Contains(selectedCommand, item) {
			continue
		}

		if item == "망치" && east.DoorType == 2 {
			east.DoorType = 0
			east.Name = "🔳"

			if len(*myItems) == 1 && (*myItems)[0] == item {
				*myItems = []string{}
			} else {
				for i := 0; i < len(*myItems); i++ {
					if (*myItems)[i] == item {
						*myItems = append((*myItems)[:i], (*myItems)[i+1])
						break
					}
				}
			}

			return true
		}
	}

	return false
}
