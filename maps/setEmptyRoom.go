package maps

import "escape-room-challenge/mapObjects"

func setEmptyRoom(room *mapObjects.Room) {
	room.DoorType = 0
	room.Name = "🔳"
}
