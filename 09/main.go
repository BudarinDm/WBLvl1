package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	chanelx1 := chanel1(array)
	chanelx2 := chanel2(chanelx1)

	for item := range chanelx2 {
		fmt.Println(item)
	}
} //наши две горутины будут почти до конца работать паралельно ,но одна будет ждать другую пока та не зальет данные
//а она уже будет ждать нашу мейн горутину пока та не считает
//получается будут работать все 3 одновременно в какой то момент

func chanel1(array []int) <-chan int {
	chanel := make(chan int) //создаем канал
	go func() {              //запускаем в горутине
		for _, item := range array { //пишем все из массива в канал
			chanel <- item
		}
		close(chanel) //закрываем тогда когда все значения запишутся
		//но писаться будет по одному потому что не буферезированный
	}()
	return chanel //возвращаем закрытый канал из которого мы сможем только читать послднее значение оставшееся там
}

func chanel2(chanel1 <-chan int) <-chan int { //тут все тоже самое
	chanel := make(chan int)
	go func() {
		for item := range chanel1 {
			chanel <- item * 2 //но тут мы отправляем в второй канал значение умноженое на2
		}
		close(chanel)
	}()
	return chanel
}
