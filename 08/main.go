package main

import (
	"fmt"
)

//Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

func main() {
	var item int64 = 20 //10100
	shiftCounter := 2
	setBit := false //true = 1, false = 0

	newItem := bitShift(item, shiftCounter, setBit)

	fmt.Println(newItem)

}

func bitShift(item int64, shiftCounter int, setBit bool) int64 {
	if setBit == true {
		newItem := item | (1 << shiftCounter) //10100 | 1000
		//побитовое ИЛИ :если хотябы в пересечении есть одна единица то ставится единица
		return newItem
	}
	newItem := item &^ (1 << shiftCounter) //10100 &^ 100
	//сброс бита (И НЕ): у newItem каждый переданый бит равен 0 ЕСЛИ бит из второго выражения равен 1,
	//если бит из второго выражения равен 0 , то берется бит из первого
	return newItem
}
