package system

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"fmt"
	"sync"
)

type Message struct {
	content string
}

// Singleton
var instance *Message
var once sync.Once

func GetMessageInstance() *Message {
	once.Do(func() {
		instance = &Message{}
	})

	return instance
}

// receivers
func (m *Message) Flush() {
	fmt.Print(m.content)
	m.content = ""
}

func (m *Message) SetCannotAct() {
	m.content += "할 수 없는 행동이거나 행동 조건을 충족시키지 못했습니다.\n"
}

func (m *Message) SetRoomIsEmpty() {
	m.content += "이 방엔 아무것도 없는 듯 하다...\n"
}

func (m *Message) SetSomethingInRoom() {
	m.content += "이 방에 무언가 있는 것 같다.\n"
}

func (m *Message) SetEmptyString() {
	m.content += ""
}

func (m *Message) SetUseItem() {
	m.content += "아이템을 사용했습니다.\n"
}

func (m *Message) SetEqpItem() {
	m.content += "아이템을 장비했습니다.\n"
}

func (m *Message) SetOpenDoor() {
	m.content += "문을 열었습니다.\n"
}

func (m *Message) SetItemsAreHere(itemType types.UsingItemTypes) {
	switch itemType {
	case types.Chest:
		m.content += "방에 상자가 있다.\n"
	case types.Hammer:
		m.content += "방에 망치가 있다.\n"
	case types.Key:
		m.content += "방에 열쇠가 있다.\n"
	case types.Potion:
		m.content += "방에 회복약이 있다.\n"
	case types.DroppedWoodSword:
		m.content += "방에 목검이 있다.\n"
	default:
	}
}

func (m *Message) SetPickUpItem(item types.UsingItemTypes) {
	switch item {
	case types.Key:
		m.content += "열쇠를 획득했다.\n"
	case types.Hammer:
		m.content += "망치를 획득했다.\n"
	case types.DroppedWoodSword:
		m.content += "목검을 획득했다.\n"
	case types.Chest:
		m.content += "상자를 획득했다.\n"
	case types.Potion:
		m.content += "회복약을 획득했다.\n"
	default:
	}
}

func (m *Message) SetCompleteRun() {
	m.content += "성공적으로 도망쳤습니다.\n"
}

func (m *Message) SetFailRun() {
	m.content += "도망에 실패했습니다. 한 대 얻어맞았습니다.\n"
}

func (m *Message) SetEnemyIsHere(enemyName string) {
	m.content += enemyName + "가(이) 있습니다.\n"
}

func (m *Message) SetGetItem(item string) {
	if item == "" {
		return
	}
	m.content += item + "을(를) 획득하셨습니다.\n"
}

func (m *Message) SetYouWin(enemy *unit.Enemy) {
	m.content += enemy.Name + "과(와)의 전투에서 승리하셨습니다.\n"
}
