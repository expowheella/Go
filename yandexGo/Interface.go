package main

import (
	"fmt"
)

// Создаем структуру Человека
type Person struct {
	Name string
	Age  int64
}

// Создаем метод для человека, который здоровается
func (p Person) sayHello() {
	fmt.Printf("Hello, My name is %s and I am %v years old", p.Name, p.Age)
}

// Создаем интерфейс, вызывающий метод sayHello
type Human interface {
	sayHello()
}

func myInterface() {
	// Создаем экземпляр человека - bu, используя интерфейс, а не структуру напрямую.
	var bu Human = Person{
		Name: "Bulat",
		Age:  34,
	}
	bu.sayHello()
}

// ---------------------------------------------------------------------------------------

// ПЕРЕДАЧА ИНТЕРФЕЙСА КАК АРГУМЕНТ ФУНКЦИИ

// создадим структуру прямоугольника
type RectangleModel struct {
	Width  int64
	Height int64
}

// создадим метод AreaMethod, которая вычисляет площадь прямоугольника
func (r RectangleModel) AreaMethod() int64 {
	return r.Width * r.Height
}

// создаем интерфейс для вызова метода расчета площади
type Shape interface {
	// есть некий интерфейс Shape, релизующий метод Area(), который возвращает float64
	AreaMethod() int64
}

// создадим функцию, где аргументом будет являться интерфейс
func CalcShape(s Shape) int64 {
	// возвращаем метод интерфейса
	return s.AreaMethod()
}

func InterfaceInArgument() {
	// создаем экземпляр структуры прямоугольника
	myRect := RectangleModel{
		Width:  40,
		Height: 20,
	}
	// т.е., мы не определили заранее переменную с типом интерфейс.
	// т.к. она будет определена при передаче в функцию CalcShape(s Shape),
	// где s - будет являться переменной с типом интерфейс.

	// вызываем функцию с передачей экземпляра структуры, как аргумент
	area := CalcShape(myRect)
	fmt.Println(area)

}
