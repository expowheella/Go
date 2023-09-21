package main

import (
	"fmt"
)

func stringExample() {
	name := "Андрей"
	firstLetter := name[0]
	fmt.Println(firstLetter) // byte: 208 (код символа "А" в таблице Unicode)

	// Преобразование строки в массив байтов
	bytes := []byte(name)
	fmt.Println(bytes)

	// Преобразование массива байтов в строку
	str := string(bytes)
	fmt.Println(str)

	// Руны
	s := "hello, мир"
	for m, r := range s {
		fmt.Printf("%c %c\n", m, r)
	}

	name2 := "異體字"
	firstLetter2 := []rune(name2)[0] // []rune(name) — конвертируем строку в слайс рун
	// и с помощью [0] забираем первый элемент
	fmt.Println(firstLetter2)         //Вывод: 30064
	fmt.Println(string(firstLetter2)) //Вывод: 異
}
