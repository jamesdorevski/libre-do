package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

var Filename string = "/Users/james.dorevski/.todo.json"

type Item struct {
	Text string
}

func SaveItems(items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(Filename, b, 0644)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ReadItems() ([]Item, error) {
	b, err := os.ReadFile(Filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
