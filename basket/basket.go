package shopshop

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// Item models an item in a shopping basket
type Item struct {
	Done  bool   `json:"done"`
	Count string `json:"count"`
	Name  string `json:"name"`
}

// Basket models the shopping list
type Basket struct {
	Color        []float64   `json:"color"`
	ShoppingList Slice[Item] `json:"shoppingList"`
	fileName     string
}

type Slice[T any] []T

func (s Slice[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte(`[]`), nil
	}
	return json.Marshal([]T(s))
}

func NewBasket() (basket *Basket) {
	basket = new(Basket)
	basket.Color = []float64{1.0, 0.84, 0, 1}
	return
}

// Save the basket at the basket's file name location in JSON format
func (sl *Basket) Save() {
	shoppingJSON, _ := json.MarshalIndent(sl, "", "	")
	err := ioutil.WriteFile(sl.fileName, shoppingJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// FileName returns the fileName where the basket is stored
func (sl *Basket) FileName() string {
	return sl.fileName
}

func (sl *Basket) ListName() string {
	return path.Base(strings.TrimSuffix(sl.FileName(), path.Ext(sl.FileName())))
}

// SetFileName set the filenName of the basket
func (sl *Basket) SetFileName(f string) {
	sl.fileName = f
}
