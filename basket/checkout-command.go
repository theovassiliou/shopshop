package shopshop

// Checkout removes all done marked items from the basket
func (sl *Basket) Checkout() {
	var newList []Item
	for _, item := range sl.ShoppingList {
		if !item.Done {
			newList = append(newList, item)
		}
	}
	sl.ShoppingList = newList

}
