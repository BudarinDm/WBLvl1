package main

import (
	"fmt"
	"math/big"
)

func main() {
	first, second := big.NewInt(3235343244322355), big.NewInt(234543242342342235) //big.NewInt конвертит наши инты
	//в бигинт

	multiplication := new(big.Int).Mul(first, second) //втстроенная в пакет функция умнножения
	division := new(big.Int).Div(first, second)       //втстроенная в пакет функция деления
	sum := new(big.Int).Add(first, second)            //втстроенная в пакет функция сложения
	subtraction := new(big.Int).Sub(first, second)    //втстроенная в пакет функция вычитания

	fmt.Println("Результат умножение:", multiplication)
	fmt.Println("Результат деления:", division)
	fmt.Println("Результат сложения:", sum)
	fmt.Println("Результат вычитания:", subtraction)
}
