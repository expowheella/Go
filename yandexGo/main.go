package main

func CheckPrime(digit int) bool {
	var count int
	for i := 1; i <= digit; i++ {
		if digit%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	// var lastDigit int
	// fmt.Scan(&lastDigit)

	// for i := 3; i <= lastDigit; i = i + 5 {
	// 	if CheckPrime(i) {
	// 		fmt.Print("хоп")
	// 	} else {
	// 		fmt.Print(i)
	// 	}
	// 	fmt.Print(" ")
	// }

	// ---------------------------------

	// Запускаем код из файла Structs.go
	// Structs()

	// Запускаем пример с прямоугольником
	// Geometry()

	// ---------------------------------

	// Запускаем Interface.go
	// myInterface()

	// Передаем интерфейс в качестве аргумента
	// InterfaceInArgument()

	// OpenFile()
	// exercise_1_1()
	Exercise_1_2()

}
