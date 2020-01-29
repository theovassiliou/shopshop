package shopshop

import (
	"bytes"
	"fmt"
)

// AddItem adds an item defined by itemDescription and with the optional quantity to the basket
func (sl *Basket) AddItem(quantity string, itemDescription []string) {
	buf := bytes.NewBuffer(nil)
	for _, word := range itemDescription[0:] {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(word)
	}
	name := buf.String()
	sl.ShoppingList = append(sl.ShoppingList, Item{Done: false, Count: quantity, Name: name})
	fmt.Println("Adding:", quantity, name)
}
