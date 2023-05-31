package maps

func RemoveItem(myItems *[]string, item string) {
	if len(*myItems) == 1 && (*myItems)[0] == item {
		*myItems = []string{}
	} else {
		newMyItems := []string{}
		for _, myItem := range *myItems {
			if myItem == item {
				continue
			}
			newMyItems = append(newMyItems, myItem)
		}
		*myItems = newMyItems
	}
}
