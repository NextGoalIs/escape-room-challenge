package maps

import (
	"escape-room-challenge/rooms"
	"regexp"
	"strings"
)

func Move(selectedCommand string, ableCommandsString string, defaultMap *[6][8]rooms.Room, nowX *int, nowY *int) bool {
	reg, _ := regexp.Compile("^동.*")

	isEast := reg.MatchString(selectedCommand)

	if isEast && strings.Contains(ableCommandsString, "동") && defaultMap[*nowX][*nowY+1].DoorType == 0 {
		setRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowY += 1
		setRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true
	}

	reg, _ = regexp.Compile("^서.*")

	isWest := reg.MatchString(selectedCommand)

	if isWest && strings.Contains(ableCommandsString, "서") && defaultMap[*nowX][*nowY-1].DoorType == 0 {
		setRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowY -= 1
		setRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true

	}

	reg, _ = regexp.Compile("^남.*")

	isSouth := reg.MatchString(selectedCommand)

	if isSouth && strings.Contains(ableCommandsString, "남") && defaultMap[*nowX-1][*nowY].DoorType == 0 {
		setRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowX -= 1
		setRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true

	}

	reg, _ = regexp.Compile("^북.*")

	isNorth := reg.MatchString(selectedCommand)

	if isNorth && strings.Contains(ableCommandsString, "북") && defaultMap[*nowX+1][*nowY].DoorType == 0 {
		setRoomNameDefaultRoom(defaultMap, *nowX, *nowY)
		*nowX += 1
		setRoomNameMyIcon(defaultMap, *nowX, *nowY)
		return true
	}

	return false
}

func setRoomNameMyIcon(defaultMap *[6][8]rooms.Room, nowX int, nowY int) {
	defaultMap[nowX][nowY].Name = "🏃"
}

func setRoomNameDefaultRoom(defaultMap *[6][8]rooms.Room, nowX int, nowY int) {
	defaultMap[nowX][nowY].Name = "🔳"
}
