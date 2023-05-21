package maps

import "escape-room-challenge/rooms"

type mapStruct struct {
	nowX int
	nowY int

	nowMap [][]rooms.Room
}

func (m mapStruct) GetThisLocation() *rooms.Room {
	return &m.nowMap[m.nowX][m.nowY]
}

func (m *mapStruct) AddX(number int) {
	m.nowX += number
}

func (m *mapStruct) AddY(number int) {
	m.nowY += number
}

func (m mapStruct) GetConnectingRooms() [4]*rooms.Room {
	var north, east, south, west *rooms.Room

	if (m.nowY+1 < len(m.nowMap[m.nowX])) && m.nowMap[m.nowX][m.nowY+1].Icon != "" {
		east = &m.nowMap[m.nowX][m.nowY+1]

	} else {
		east = rooms.GetWall()
	}

	if (m.nowY-1 >= 0) && m.nowMap[m.nowX][m.nowY-1].Icon != "" {
		west = &m.nowMap[m.nowX][m.nowY-1]

	} else {
		west = rooms.GetWall()
	}

	if (m.nowX-1 >= 0) && m.nowMap[m.nowX-1][m.nowY].Icon != "" {
		south = &m.nowMap[m.nowX-1][m.nowY]

	} else {
		south = rooms.GetWall()
	}

	if (m.nowX+1 < len(m.nowMap) && m.nowY < len(m.nowMap[m.nowX+1])) && m.nowMap[m.nowX+1][m.nowY].Icon != "" {
		north = &m.nowMap[m.nowX+1][m.nowY]
	} else {
		north = rooms.GetWall()
	}
	return [4]*rooms.Room{east, west, south, north}
}
