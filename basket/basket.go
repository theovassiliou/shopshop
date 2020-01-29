package shopshop

import (
	"encoding/json"
	"io/ioutil"
)

// Item models an item in a shopping basket
type Item struct {
	Done  bool   `json:"done"`
	Count string `json:"count"`
	Name  string `json:"name"`
}

// Basket models the shopping list
type Basket struct {
	Color        []float64 `json:"color"`
	ShoppingList []Item    `json:"shoppingList"`
	fileName     string
}

// Save the basket at the basket's file name location in JSON format
func (sl *Basket) Save() {
	shoppingJSON, _ := json.MarshalIndent(sl, "", "  ")
	err := ioutil.WriteFile(sl.fileName, shoppingJSON, 0644)
	AssertNoError(err)
}

// FileName returns the fileName where the basket is stored
func (sl *Basket) FileName() string {
	return sl.fileName
}

// SetFileName set the filenName of the basket
func (sl *Basket) SetFileName(f string) {
	sl.fileName = f
}
