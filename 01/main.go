package main

import "fmt"

type Human struct {
	name string
	age  int
	work bool
}

type Action struct {
	Human
}

func (h *Human) SetWork(work bool) {
	h.work = work
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) GetName() {
	fmt.Println("Человек с именем", h.name)
}

func (h *Human) GetAge() {
	fmt.Println("Человек в возрасте", h.age)
}

func (h *Human) GetWork() {
	if h.work {
		fmt.Println("Работает")
		return
	}
	fmt.Println("Не работает")
}

func main() {
	action := new(Action)

	action.SetName("Крутое имя")
	action.SetAge(23)
	action.SetWork(true)
	action.GetName()
	action.GetAge()
	action.GetWork()
}
