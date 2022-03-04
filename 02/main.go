package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	squaring(array)

}

func squaring(array []int) {
	var wg sync.WaitGroup //создается для групировки горутин, и ждет пока они не выполнятся

	for _, item := range array {
		wg.Add(1) //устанавливаем сколько горутин должно выполнится в группе, каждую итерацию добаляем 1
		go pow(&wg, item)
	}

	//Либо можно так:
	//wg.Add(len(array))
	//for _, item := range array {
	//	go pow(&wg, item)
	//}

	wg.Wait() //закрываем группу, пока группа открыта дальше двигатьcя процесс не будет
	//как я понимаю нам нужно свести счетчик, который ставим Add к 0?? и как я понимаю если будет запущено больше горутин
	//чем счетчик, то 100% закончат работать столько горутин сколько указали в счетчике , а остальные могут не успеть
	//завершится
}

func pow(wg *sync.WaitGroup, item int) {
	defer wg.Done() //здесь же мы убавляем счетчик в группе на 1
	fmt.Println(item * item)
}
