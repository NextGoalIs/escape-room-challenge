package system

func AddUseItemCommands(myItems []string, ableCommands *[]string) {
	for _, v := range myItems {
		*ableCommands = append(*ableCommands, v+" 사용")
	}
}
