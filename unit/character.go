package unit

import (
	"escape-room-challenge/utils"
	"fmt"
)

type Character struct {
	Name string
	Icon string

	Items []string
}

func (c *Character) SetName() {
	for c.Name == "" {
		fmt.Println("캐릭터 이름을 1~16자 사이로 입력해주세요.")
		fmt.Print(">>> ")
		fmt.Scanln(&c.Name)
		if len(c.Name) > 16 || c.Name == "" {
			utils.ClearConsoleWindows()
			fmt.Println("캐릭터 이름이 잘못되었습니다.")
			c.Name = ""
		}
	}
}

func (c *Character) LookAt(commnad string) bool {
	switch commnad {
	case "방":
		return true
	default:
		return false
	}
}
