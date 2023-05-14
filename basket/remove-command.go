package shopshop

import (
	"fmt"
	"strconv"
)

func isInIndices(ref int, indices []string) bool {
	for _, index := range indices {
		idx, err := strconv.Atoi(index)
		if err == nil {
			// continue here on no error
			if ref == idx {
				return true
			}
		}
	}
	return false
}

// Remove removes the indicated list items from the basket
func (sl *Basket) Remove(Indices []string) {
	var newList []Item
	for i, shopItem := range sl.ShoppingList {
		if !isInIndices(i, Indices) {
			newList = append(newList, shopItem)
		} else {
			fmt.Println("Removing:", shopItem.Count, shopItem.Name)
		}
	}
	sl.ShoppingList = newList
}
