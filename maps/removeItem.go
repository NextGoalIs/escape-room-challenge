package maps

func RemoveItem(myItems *[]string, item string) {
	if len(*myItems) == 1 && (*myItems)[0] == item {
		*myItems = []string{}
	} else {
		for i, myItem := range *myItems {
			if myItem == item {
				*myItems = append((*myItems)[:i], (*myItems)[i+1:]...)
				break
			}
		}
	}
}
