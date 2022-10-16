package main

import "fmt"

func main() {
	// динамический неприрывный массив

	// создаем пустой список для чисел. Вариант 1
	// var vector []int

	// создаем пустой список для чисел. Вариант 2
	vector := []int{}
	fmt.Println(vector)

	// стек
	stack := vector

	// PUSH. добавление в стек при помощи append (добавляем в конец)
	// stack = append (stack, vlaue)
	for value := 0; value < 10; value++ {
		stack = append(stack, value)
		fmt.Println(stack)
	}

	// POP. изъятие из стека (забираем с конца)
	fmt.Println("изъятие из стека")

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("popped item is", item, stack)

	}

}
