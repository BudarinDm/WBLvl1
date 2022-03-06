package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(4.235, 7.235)
	p2 := NewPoint(5.342, 3.325)

	fmt.Println(distance(p1, p2))
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

//AB = √((xb - xa)^2 + (yb - ya)^2)

func distance(p1, p2 *Point) float64 {
	sumPow := math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2) //pow возводит в степень которую мы передаем
	//вторым аргрументом
	return math.Sqrt(sumPow) //извлекает корень
}
