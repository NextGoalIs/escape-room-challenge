package rooms

import "escape-room-challenge/types"

func (room *Room) SetMyCharacter() {
	room.Icon = string(types.MyCharacterIcon)
}
