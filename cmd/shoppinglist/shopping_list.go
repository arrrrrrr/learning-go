package main

import "fmt"

type shopping_list_item struct {
	name   string
	amount int
}

func add_to_shopping_list(
	shopping_list []shopping_list_item,
	name string,
	amount int,
) []shopping_list_item {

	shopping_list = append(shopping_list, shopping_list_item{name, amount})
	return shopping_list
}

func main() {
	shopping_list := []shopping_list_item{}
	shopping_list = add_to_shopping_list(shopping_list, "Tissues", 2)
	shopping_list = add_to_shopping_list(shopping_list, "Tim Tams", 1)
	shopping_list = add_to_shopping_list(shopping_list, "Banana", 4)

	for _, v := range shopping_list {
		fmt.Printf("Item: %s, Number: %d\n", v.name, v.amount)
	}
}
