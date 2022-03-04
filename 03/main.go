package main

import "fmt"

func main() {
	array := []int{3, 5, 3}
	chanel := make(chan int, len(array)) //создаем буферизированый канал который будет хранить в себе столько значений,
	//сколько обьектов в нашем слайсе
	SumSquaring(chanel, array)

}

func SumSquaring(chanel chan int, array []int) {
	for _, item := range array { //бежим по нашему массиву достаем итемы
		go squaring(chanel, item) //и запускаем несколько паралельных функций , которые накидают числа в квадрате в канал
	}

	var sum int
	for i := 0; i < len(array); i++ {
		sum += <-chanel //достает все накиданные значения и сумирует
	}

	fmt.Println(sum)
}

func squaring(chanel chan int, item int) {
	chanel <- item * item //кидает в канал число в квадрате
}
