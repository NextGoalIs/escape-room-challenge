package rooms

import "escape-room-challenge/types"

func (room *Room) SetWoodDoor() {
	room.DoorType = types.Wood
	room.Name = "🚪"
}
