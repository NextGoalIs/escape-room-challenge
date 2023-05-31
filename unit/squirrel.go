package unit

func NewSquirrel() Enemy {
	e := Enemy{Name: "다람쥐", Health: 50, AttackDamage: 5, Defence: 0, dropTable: map[string]int{"회복약": 70}}

	return e
}
