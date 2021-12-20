package main

import "fmt"

type shopping_list_item struct {
	name   string
	amount int
}

func main() {
	fmt.Println("This is a shopping list")
	shopping_list := []shopping_list_item{}
	shopping_list = append(shopping_list, shopping_list_item{"corn", 1})

	for _, v := range shopping_list {
		fmt.Printf("Item: %s, Number: %d\n", v.name, v.amount)
	}
}
