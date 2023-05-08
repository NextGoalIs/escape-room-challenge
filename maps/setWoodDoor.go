package maps

import "escape-room-challenge/mapObjects"

func setWoodDoor(room *mapObjects.Room) {
	room.DoorType = 1
	room.Name = "🚪"
}
