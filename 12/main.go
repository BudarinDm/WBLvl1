package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	items := []string{"cat", "cat", "dog", "cat", "tree"}

	newItems := uniqItems(items)

	fmt.Println(newItems)
}

func uniqItems(items []string) []string {
	uniqItemsMap := map[string]int{}

	for _, item := range items { //кидаем в ключи наши значения из слайса, в мапе не бывает одинаковых ключей,они все
		//уникальны
		uniqItemsMap[item] = 0
	}

	var newItems []string

	for key, _ := range uniqItemsMap { //добавляем ключи в слайс игнорируя значения по ключам
		newItems = append(newItems, key)
	}
	return newItems
}
