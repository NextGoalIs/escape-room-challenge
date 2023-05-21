package unit

import (
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
)

type Character struct {
	Name string
	Icon string

	health       int
	attackDamage int
	defence      int

	Items        []string
	EquipedItems EqpItems
}

type EqpItems struct {
	Shirt     types.ShirtItemTypes
	Pants     types.PantsItemTypes
	Shoes     types.ShoesItemTypes
	LeftHand  types.WeaponTypes
	RightHand types.WeaponTypes
}

func NewCharacter() Character {
	c := Character{}
	c.Icon = string(types.MyCharacterIcon)

	c.health = 50
	c.attackDamage = 3
	c.defence = 0

	return c
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

func (c *Character) ShowStatus() {

}
