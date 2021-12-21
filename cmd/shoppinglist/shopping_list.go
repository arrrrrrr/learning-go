package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type shoppingListItem struct {
	amount int
}

type ShoppingListContainer map[string]shoppingListItem

func addToShoppingList(
	shoppingList ShoppingListContainer,
	name string,
	amount int,
) {
	updateShoppingList(shoppingList, name, shoppingListItem{amount})
}

func removeFromShoppingList(
	shoppingList ShoppingListContainer,
	name string,
	amount int,
) {
	updateShoppingList(shoppingList, name, shoppingListItem{-amount})
}

func updateShoppingList(
	shoppingList ShoppingListContainer,
	name string,
	item shoppingListItem,
) {
	if oldItem, ok := shoppingList[name]; ok {
		item.amount += oldItem.amount
	}
	if shoppingList[name] = item; shoppingList[name].amount <= 0 {
		delete(shoppingList, name)
	}
}

func printShoppingList(
	shoppingList ShoppingListContainer,
	itemNames []string,
) {
	fmt.Printf("Shopping List\n")
	sort.Strings(itemNames)

	for _, k := range itemNames {
		if v, ok := shoppingList[k]; ok {
			fmt.Printf("  %s\t[%d]\n", k, v.amount)
		}
	}
}

func readAmount(
	scanner *bufio.Scanner,
) (int, error) {
	if scanner.Scan() {
		return strconv.Atoi(scanner.Text())
	}
	return 0, scanner.Err()
}

func readAddAmount(
	scanner *bufio.Scanner,
	name string,
) (int, error) {
	fmt.Printf("Amount of %s to add? ", name)
	return readAmount(scanner)
}

func readRemoveAmount(
	scanner *bufio.Scanner,
	name string,
) (int, error) {
	fmt.Printf("Amount of %s to remove? ", name)
	return readAmount(scanner)
}

func handleModifyCommand(
	cmd string,
	data string,
	scanner *bufio.Scanner,
	shoppingList ShoppingListContainer,
) {
	var funcReadAmount func(*bufio.Scanner, string) (int, error)
	var funcUpdateList func(ShoppingListContainer, string, int)

	if cmd == "add" {
		funcReadAmount = readAddAmount
		funcUpdateList = addToShoppingList
	} else if cmd == "remove" {
		funcReadAmount = readRemoveAmount
		funcUpdateList = removeFromShoppingList
	}

	for {
		if amount, err := funcReadAmount(scanner, data); err != nil || amount <= 0 {
			var errStr string
			if amount <= 0 {
				errStr = "Amount must be greater than zero"
			} else {
				errStr = err.Error()
			}
			fmt.Printf("ERROR: %s\n", errStr)
		} else {
			funcUpdateList(shoppingList, data, amount)
			return
		}
	}
}

func handleShowCommand(
	shoppingList ShoppingListContainer,
) {
	keys := make([]string, 0)
	for k := range shoppingList {
		keys = append(keys, k)
	}
	printShoppingList(shoppingList, keys)
}

func handleInvalidCommand() {
	fmt.Printf("HELP: add <item name>|remove <item name>|show|quit")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	shoppingList := make(ShoppingListContainer, 0)

	for {
		fmt.Printf(">: ")
		if !scanner.Scan() {
			fmt.Printf("FATAL: %s\n", scanner.Err())
			return
		}

		line := scanner.Text()
		tokens := strings.SplitN(line, " ", 2)
		cmd := tokens[0]

		switch cmd {
		case "add", "remove":
			if len(tokens) != 2 {
				handleInvalidCommand()
			} else {
				handleModifyCommand(cmd, tokens[1], scanner, shoppingList)
			}
		case "show":
			handleShowCommand(shoppingList)
		case "quit":
			return
		default:
			handleInvalidCommand()
		}

		fmt.Println()
	}
}
