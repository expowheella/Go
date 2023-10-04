package main

import (
	"fmt"
)

// Структура Студент
type Student struct {
	Name string
	Age  int
}

// Метод должен быть связан с определённой структурой.
// Перед именем функции указывается имя структуры (так называемый ресивер), с которой она связана:
func (s Student) printData() {
	fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
}

// Метод printData связан со структурой Student.
// Метод принимает один аргумент типа Student, который обозначается буквой s (можно обозначить любым названием).
// Внутри метода мы можем обращаться к полям структуры с помощью обычного синтаксиса s.<имя поля>.

// Вызов метода:

func Structs() {
	student1 := Student{Name: "vasya", Age: 15}
	student1.printData()
}

// --------------------------------------------------------------------------------------------------------------
// Пример Geometry

// Структура Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// метод вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// метод в качестве аргумента х получает множитель
// метод возрващается новые размеры прямоугольника (ширину, высоту)
func (r Rectangle) Scale(x int) (float64, float64) {
	r.Width *= float64(x)
	r.Height *= float64(x)
	return r.Width, r.Height
}

// Программа
func Geometry() {
	// создадим экземпляр прямоугольника
	rectangle := Rectangle{
		Width:  2.0,
		Height: 3.0,
	}

	// \n - new line
	fmt.Printf("Ширина: %v, Высота: %v \n", rectangle.Width, rectangle.Height)

	// вызовим метод для вычисления площади прямоугольника
	area := rectangle.Area()
	fmt.Printf("Площадь: %v \n", area)

	// вызовим метод для увеличения прямоугольника в 5 раз
	width, height := rectangle.Scale(5)
	fmt.Printf("Ширина: %v, Высота: %v", width, height)

}
