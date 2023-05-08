package rooms

import "escape-room-challenge/types"

func (room *Room) SetMyCharacterRoom() {
	room.Name = string(types.MyCharacterName)
}
