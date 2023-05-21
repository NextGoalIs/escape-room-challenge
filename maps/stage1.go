package maps

import (
	"escape-room-challenge/rooms"
	"escape-room-challenge/types"
)

func GetStage1() mapStruct {
	var defaultMap = setRooms(6, 8)
	defaultMap[3][0] = rooms.NewRoom(types.NoDoorType, types.Hammer, false) // 망치
	defaultMap[1][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false) // 스타트지점
	defaultMap[2][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][1] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][2] = rooms.NewRoom(types.GlassType, types.NoItem, false) // 유리문 방
	defaultMap[1][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[2][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[3][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[4][3] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[0][4] = rooms.NewRoom(types.NoDoorType, types.Key, false)  // 열쇠
	defaultMap[1][4] = rooms.NewRoom(types.WoodType, types.NoItem, false) // 나무문
	defaultMap[4][4] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][4] = rooms.NewRoom(types.WoodType, types.NoItem, false) // 나무문
	defaultMap[5][5] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][6] = rooms.NewRoom(types.NoDoorType, types.NoItem, false)
	defaultMap[5][7] = rooms.NewRoom(types.LockedType, types.NoItem, true) // 열쇠가 필요한 문

	defaultMap[1][1].SetMyCharacter()

	return mapStruct{nowX: 1, nowY: 1, nowMap: defaultMap}
}

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
