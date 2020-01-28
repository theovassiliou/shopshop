package shopshop

type Item struct {
	Done  bool   `json:"done"`
	Count string `json:"count"`
	Name  string `json:"name"`
}

type Basket struct {
	Color        []float64 `json:"color"`
	ShoppingList []Item    `json:"shoppingList"`
}
