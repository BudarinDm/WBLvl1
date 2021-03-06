package main

import "fmt"

func main() {
	array := []int{1, 5, 4, 46, 5, 2, 3, 73, 3, 5, 7, 9, 32, 76, 83, 2, 65, 8, 2, 4, 53, 67, 16}

	arrSort := quicksort(array, 0, len(array)-1)
	fmt.Println(arrSort)
}

func quicksort(arr []int, first, last int) []int {
	l, r := first, last //делаем копии наших переданных крайних значений//мы можем сортировать определенный кусок массива
	piv := arr[(l+r)/2] //находим наше опорное значение от которого будем отталкиваться

	for l <= r {
		for arr[l] < piv { //идем с крайнего левого элемента пока оно не окажется больше опорного
			l++
		}
		for arr[r] > piv { //идем от крайнего правого элемента пока оно не окажется меньше опорного
			r--
		}

		if l <= r { //если l все еще меньше r то меняем местами обьекты по индексам l и r
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if first < r { //если наш r больше первого заданого индекса в массиве то
		quicksort(arr, first, r) //делаем рекурсию но уже  берем часть массива с первого элемента по элемент по индексу r
	}
	if last > l { //тоже самое но если последний заданный элемент больше левого индекса
		quicksort(arr, l, last) // так же рекурсия с элемента на котором остановилсь по индексу l до последнего заданого
	}
	return arr
}
