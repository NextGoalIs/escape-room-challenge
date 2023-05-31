package maps

import (
	"escape-room-challenge/types"
	"escape-room-challenge/unit"
	"escape-room-challenge/utils"
)

func EqpItem(firstCommand string, char *unit.Character) bool {
	if !utils.Contains(char.Items, firstCommand) {
		return false
	}

	for weaponType, name := range types.WeaponNameMap {
		if name != firstCommand {
			continue
		}

		switch types.NoWeaponItem {
		case int(char.LeftHand):
			char.LeftHand = weaponType
			RemoveItem(&char.Items, name)
			return true
		case int(char.RightHand):
			char.RightHand = weaponType
			RemoveItem(&char.Items, name)
			return true
		default:
			return false
		}
	}

	for shoesItem, name := range types.ShoesNameMap {
		if name != firstCommand {
			continue
		}

		switch types.NoShoesItem {
		case char.Shoes:
			char.Shoes = shoesItem
			RemoveItem(&char.Items, name)
			return true
		default:
			return false
		}
	}

	for shirtItem, name := range types.ShirtNameMap {
		if name != firstCommand {
			continue
		}

		switch types.NoShirtItem {
		case char.Shirt:
			char.Shirt = shirtItem
			RemoveItem(&char.Items, name)
			return true
		default:
			return false
		}
	}

	for pantsItem, name := range types.PantsNameMap {
		if name != firstCommand {
			continue
		}

		switch types.NoPantsItem {
		case char.Pants:
			char.Pants = pantsItem
			RemoveItem(&char.Items, name)
			return true
		default:
			return false
		}
	}

	return false
}
