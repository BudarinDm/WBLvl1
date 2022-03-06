package main

import "fmt"

func main() {
	arr0 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr0 = deleteByIndex0(arr0, 5)
	fmt.Println(arr0)
	arr1 = deleteByIndex1(arr1, 3)
	fmt.Println(arr1)
	arr2 = deleteByIndex2(arr2, 7)
	fmt.Println(arr2)
}

func deleteByIndex0(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func deleteByIndex1(arr []int, i int) []int {
	//Копировать последний элемент в индекс i.
	arr[i] = arr[len(arr)-1]

	//Удалить последний элемент (записать нулевое значение).
	arr[len(arr)-1] = 0

	//Усечь срез.
	arr = arr[:len(arr)-1]

	return arr
}

func deleteByIndex2(arr []int, i int) []int {
	//Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(arr[i:], arr[i+1:])

	//Удалить последний элемент (записать нулевое значение).
	arr[len(arr)-1] = 0

	//Усечь срез.
	arr = arr[:len(arr)-1]

	return arr
}
