package rooms

import "escape-room-challenge/types"

func (room *Room) SetMyCharacterRoom() {
	room.Icon = string(types.MyCharacterIcon)
}
