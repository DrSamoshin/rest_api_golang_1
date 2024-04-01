package main

import (
	"fmt"
	"math"
)

// объединение интерфейсов
type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Area() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}
	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(shape ShapeWithArea) {
	fmt.Println(shape.Area())
}
