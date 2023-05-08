package maps

import "escape-room-challenge/rooms"

func setWoodDoor(room *rooms.Room) {
	room.DoorType = 1
	room.Name = "🚪"
}
