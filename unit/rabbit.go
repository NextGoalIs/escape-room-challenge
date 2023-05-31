package unit

func NewRabbit() Enemy {
	e := Enemy{Name: "토끼", Health: 70, AttackDamage: 7, Defence: 3, dropTable: map[string]int{"회복약": 50, "회복약회복약": 30}}

	return e
}
