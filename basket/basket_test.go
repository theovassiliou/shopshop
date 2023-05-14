package shopshop

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewBasket(t *testing.T) {
	basket := NewBasket()

	// Assert the default color values
	expectedColor := []float64{1.0, 0.84, 0, 1}
	if !reflect.DeepEqual(basket.Color, expectedColor) {
		t.Errorf("Unexpected color value. Expected: %v, Got: %v", expectedColor, basket.Color)
	}

	// Assert the shopping list is empty
	if len(basket.ShoppingList) != 0 {
		t.Errorf("Unexpected shopping list length. Expected: 0, Got: %d", len(basket.ShoppingList))
	}
}

func TestSave(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "test_basket_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Set the temporary file name for the basket
	basket := NewBasket()
	basket.SetFileName(tmpFile.Name())

	// Save the basket
	basket.Save()

	// Read the content of the saved file
	fileContent, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file content: %v", err)
	}

	// Assert the content matches the expected JSON format
	expectedContent := `{
	"color": [
		1,
		0.84,
		0,
		1
	],
	"shoppingList": []
}`
	if string(fileContent) != expectedContent {
		t.Errorf("Unexpected file content. Expected:\n%s\nGot:\n%s", expectedContent, string(fileContent))
	}
}

func TestFileNameAndListName(t *testing.T) {
	// Create a basket and set the file name
	basket := NewBasket()
	fileName := "/path/to/basket.json"
	basket.SetFileName(fileName)

	// Assert the file name
	if basket.FileName() != fileName {
		t.Errorf("Unexpected file name. Expected: %s, Got: %s", fileName, basket.FileName())
	}

	// Assert the list name (without file extension)
	expectedListName := "basket"
	if basket.ListName() != expectedListName {
		t.Errorf("Unexpected list name. Expected: %s, Got: %s", expectedListName, basket.ListName())
	}
}
