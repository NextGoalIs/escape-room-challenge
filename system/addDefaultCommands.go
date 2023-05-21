package system

func AddDefaultCommands(ableCommands *[]string) {
	*ableCommands = append(*ableCommands, "<목표> 보기")
	*ableCommands = append(*ableCommands, "<목표> 줍기")
}
