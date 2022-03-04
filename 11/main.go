package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{4, 5, 0}

	is := intersection(array1, array2)
	fmt.Println(is)
}

func intersection(array1, array2 []int) []int {
	isMap := map[int]bool{} //мап для хранения уникальных

	for _, item1 := range array1 { //бежим каждым элементом первого массива
		for _, item2 := range array2 { //по каждому элементу второго массива
			if item1 == item2 { //если значение из первого массива равно значению второго
				isMap[item1] = true //то пишем в мапу тру, почему мапа а не слайс? потому что мы избавляемся от дублирования
			} //если в одном из масивов несколько одиноковых элементов , то в мапу перезапишется по новой с тру
		}
	}

	var is []int

	for itemMap := range isMap { //просто переписываем из ключей мапы значения в слайс
		is = append(is, itemMap)
	}
	return is //возвращаем слайс
}
