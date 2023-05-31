package unit

import (
	"escape-room-challenge/types"
	"escape-room-challenge/utils"
	"fmt"
	"strings"
)

type Character struct {
	Name string
	Icon string

	Health       int
	AttackDamage int
	Defence      int

	Items        []string
	EqpShirt     types.ShirtItemTypes
	EqpPants     types.PantsItemTypes
	EqpShoes     types.ShoesItemTypes
	EqpLeftHand  types.WeaponTypes
	EqpRightHand types.WeaponTypes
}

func NewCharacter() Character {
	c := Character{}
	c.Icon = string(types.MyCharacterIcon)

	c.Health = 50
	c.AttackDamage = 3
	c.Defence = 0

	c.SetName()

	return c
}

func (c Character) getCalcAttackDamage() int {
	return c.AttackDamage + types.WeaponDamageMap[c.EqpLeftHand] + types.WeaponDamageMap[c.EqpRightHand]
}

// TODO 미완
// func (c Character) getCalcAttackDamageDetail() string {
// 	content := string(c.getCalcAttackDamage())


// 	return content
// }

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

func (c Character) AttackTo(enemy *Enemy) {
	damage := enemy.Defence - c.getCalcAttackDamage()
	if damage >= 0 {
		return
	}
	enemy.Health = enemy.Health + damage
}

func (c *Character) ShowStatus() {
	utils.ClearConsoleWindows()
	fmt.Println("캐릭터 이름 : ", c.Name)
	fmt.Println("장비한 아이템 : ", "미완성 ㅎㅎ;")
	fmt.Println("체력 : ", c.Health)
	fmt.Println("공격력 : ", ) //TODO getCalcAttackDamageDetail로 변경하기
	fmt.Println("방어력 : ", c.Defence)
	fmt.Println("소유한 아이템 : ", strings.Join(c.Items, ", "))
	fmt.Println()
	fmt.Println("돌아가려면 Enter키를 눌러주세요")
	fmt.Scanln()
}
