package maps

import (
	"regexp"
	"strings"
)

func Move(selectedCommand string, ableCommandsString string, stage *mapStruct) bool {
	reg, _ := regexp.Compile("^동.*")

	isEast := reg.MatchString(selectedCommand)

	if isEast && strings.Contains(ableCommandsString, "동") && stage.nowMap[stage.nowX][stage.nowY+1].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(1)
		stage.GetThisLocation().SetMyCharacter()
		return true
	}

	reg, _ = regexp.Compile("^서.*")

	isWest := reg.MatchString(selectedCommand)

	if isWest && strings.Contains(ableCommandsString, "서") && stage.nowMap[stage.nowX][stage.nowY-1].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(-1)
		stage.GetThisLocation().SetMyCharacter()
		return true

	}

	reg, _ = regexp.Compile("^남.*")

	isSouth := reg.MatchString(selectedCommand)

	if isSouth && strings.Contains(ableCommandsString, "남") && stage.nowMap[stage.nowX-1][stage.nowY].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(-1)
		stage.GetThisLocation().SetMyCharacter()
		return true

	}

	reg, _ = regexp.Compile("^북.*")

	isNorth := reg.MatchString(selectedCommand)

	if isNorth && strings.Contains(ableCommandsString, "북") && stage.nowMap[stage.nowX+1][stage.nowY].DoorType == 0 {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(1)
		stage.GetThisLocation().SetMyCharacter()
		return true
	}

	return false
}
