package maps

import "escape-room-challenge/rooms"

func GetMap() [6][8]rooms.Room {
	var defaultMap [6][8]rooms.Room = [6][8]rooms.Room{}
	defaultMap[3][0] = rooms.NewRoom(0, 2, false) // 망치
	defaultMap[1][1] = rooms.NewRoom(0, 0, false) // 스타트지점
	defaultMap[2][1] = rooms.NewRoom(0, 0, false)
	defaultMap[3][1] = rooms.NewRoom(0, 0, false)
	defaultMap[3][2] = rooms.NewRoom(2, 0, false) // 유리문 방
	defaultMap[1][3] = rooms.NewRoom(0, 0, false)
	defaultMap[2][3] = rooms.NewRoom(0, 0, false)
	defaultMap[3][3] = rooms.NewRoom(0, 0, false)
	defaultMap[4][3] = rooms.NewRoom(0, 0, false)
	defaultMap[0][4] = rooms.NewRoom(0, 1, false) // 열쇠
	defaultMap[1][4] = rooms.NewRoom(1, 0, false) // 나무문
	defaultMap[4][4] = rooms.NewRoom(0, 0, false)
	defaultMap[5][4] = rooms.NewRoom(1, 0, false) // 나무문
	defaultMap[5][5] = rooms.NewRoom(0, 0, false)
	defaultMap[5][6] = rooms.NewRoom(0, 0, false)
	defaultMap[5][7] = rooms.NewRoom(3, 0, true) // 열쇠가 필요한 문

	return defaultMap
}
