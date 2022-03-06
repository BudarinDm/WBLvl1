package main

import "fmt"

func main() {
	text := "qQwWert"

	fmt.Println(uniq(text))
}

func uniq(text string) bool {
	symbolArr := make(map[rune]bool) //мапа с ключем руной и значением тру если такой элемент уже был
	for _, symbol := range text {
		_, ok := symbolArr[symbol] //ищем значение по ключю, если оно уже было в мапе то ретрн фалс потому что уже
		if ok {                    //не уникальные
			return false
		}
		symbolArr[symbol] = true //тут пишем мапу значение если его еще не было
	}

	return true
}
