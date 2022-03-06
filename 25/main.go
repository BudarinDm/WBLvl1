package main

import (
	"fmt"
	"time"
)

func main() {
	customTimer(3)
}

//After ожидает истечения продолжительности d, а затем отправляет текущее время по возвращаемому каналу.

func customTimer(second int) {
	//<-time.After(time.Duration(second) * time.Second)
	fmt.Println(<-time.After(time.Duration(second) * time.Second))
}
