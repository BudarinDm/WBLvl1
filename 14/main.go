package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeSpecifier(true)
	typeSpecifier1(43)
	typeSpecifier2("dfs")
}

func typeSpecifier(param interface{}) {
	fmt.Println(reflect.TypeOf(param))
}

func typeSpecifier1(param interface{}) {
	fmt.Printf("%T\n", param)
}

func typeSpecifier2(param interface{}) { //ну тут без каналов , на них можно сделать кейсы еще на каждый тип
	switch t := param.(type) {
	case int:
		fmt.Println(t)
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	default:
		fmt.Println("unknown")
	}
}
