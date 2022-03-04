package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var workerCount int
	fmt.Println("Укажите количество воркеров:")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		log.Fatal("Вы ввели не число")
	}

	chanel := make(chan interface{})                     //канал для воркеров который принимает любые данные
	done := make(chan os.Signal, 1)                      //буферезированный канал который принимает один сигнал от ОС
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM) //тут мы ловим и записываем сигнал от ОС в канал

	for i := 0; i < workerCount; i++ {
		go worker(chanel) //запускам столько воркеров сколько указал пользователь
	}

	for i := 0; ; i++ {
		select { // оператор итерации по каналам. причем будет исполнен каждый кейс в отличии от свича
		case <-done: //читаем канал done и если туда что-то приходит то
			close(chanel) //закрываем канал для воркеров
			close(done)   //закрываем канал для чтение сигналов ос
			return        //выходим из бесконечного цикла и основаня горутина мейн тоже завершится
		default:
			if i%2 == 0 {
				chanel <- fmt.Sprintf("text %d", i) //пишем в канал воркеров
				time.Sleep(2 * time.Second)
				continue
			}
			chanel <- i // тоже пишем
			time.Sleep(2 * time.Second)
		}
	}

}

func worker(chanel <-chan interface{}) { //принимает канал доступный тут только для чтения
	for item := range chanel { //бежим циклом по каналу доставая все что сюда приходит,
		fmt.Println(item) // как я понимаю цикл будет бесконечно читать из канала пока он не закроется и потом
	} //выход из цикла и горутина завершится
}
