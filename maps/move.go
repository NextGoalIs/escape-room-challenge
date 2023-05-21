package maps

func Move(selectedCommand string, stage *MapStruct) bool {

	connectingRooms := stage.GetConnectingRooms()
	north := connectingRooms[3]
	west := connectingRooms[1]
	east := connectingRooms[0]
	south := connectingRooms[2]

	if selectedCommand == "동" && stage.CanMove(east) {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(1)
		stage.GetThisLocation().SetMyCharacter()

		return true
	}

	if selectedCommand == "서" && stage.CanMove(west) {
		stage.GetThisLocation().SetEmpty()
		stage.AddY(-1)
		stage.GetThisLocation().SetMyCharacter()

		return true
	}

	if selectedCommand == "남" && stage.CanMove(south) {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(-1)
		stage.GetThisLocation().SetMyCharacter()

		return true
	}

	if selectedCommand == "북" && stage.CanMove(north) {
		stage.GetThisLocation().SetEmpty()
		stage.AddX(1)
		stage.GetThisLocation().SetMyCharacter()

		return true
	}

	return false
}
