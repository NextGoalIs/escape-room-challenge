package maps

import (
	"escape-room-challenge/utils"
	"regexp"
)

func Move(selectedCommand string, ableCommands []string, stage *MapStruct) bool {
	reg, _ := regexp.Compile("^동.*")

	isEast := reg.MatchString(selectedCommand)

	if isEast && utils.Contains(ableCommands, "동") && stage.nowMap[stage.nowX][stage.nowY+1].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(1)
		stage.GetThisLocation().SetMyCharacter()
		return true
	}

	reg, _ = regexp.Compile("^서.*")

	isWest := reg.MatchString(selectedCommand)

	if isWest && utils.Contains(ableCommands, "서") && stage.nowMap[stage.nowX][stage.nowY-1].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(-1)
		stage.GetThisLocation().SetMyCharacter()
		return true

	}

	reg, _ = regexp.Compile("^남.*")

	isSouth := reg.MatchString(selectedCommand)

	if isSouth && utils.Contains(ableCommands, "남") && stage.nowMap[stage.nowX-1][stage.nowY].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(-1)
		stage.GetThisLocation().SetMyCharacter()
		return true

	}

	reg, _ = regexp.Compile("^북.*")

	isNorth := reg.MatchString(selectedCommand)

	if isNorth && utils.Contains(ableCommands, "북") && stage.nowMap[stage.nowX+1][stage.nowY].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(1)
		stage.GetThisLocation().SetMyCharacter()
		return true
	}

	return false
}
