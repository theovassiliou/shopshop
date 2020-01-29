package shopshop

import "fmt"

// List all items of the basket to stdout
func (sl *Basket) List() {
	fmt.Println("Items:")
	for i, item := range sl.ShoppingList {
		check := " "
		if item.Done {
			check = "@done"
		}
		fmt.Printf("%2d: %s %s %s\n", i, item.Count, item.Name, check)
	}
	fmt.Println()
}
