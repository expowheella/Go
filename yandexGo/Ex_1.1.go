// 1 сентября каждого учебного года во всех базах данных школьников происходит великий пересчёт.
// Напишите функцию modifyJSON(jsonData []byte) ([]byte, error), которая принимает данные в формате JSON,
// добавляет 1 год к полю grade и возвращает обновлённые данные также в формате JSON.

// Формат ввода
// [
//     {
//         "name": "Oleg",
//         "grade": 9
//     },
//     {
//         "name": "Katya",
//         "grade": 10
//     }
// ]
// Формат вывода
// [
//     {
//         "name": "Oleg",
//         "grade": 10
//     },
//     {
//         "name": "Katya",
//         "grade": 11
//     }
// ]
// Примечания
// Структура ученика

//	type Student struct {
//	    Name  string `json:"name"`
//	    Grade int    `json:"grade"`
//	}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student_ struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {

	var students []Student_

	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		fmt.Print("Unmarshaling error:", err)
	}

	for i := 0; i < len(students); i++ {
		students[i].Grade++
	}

	jsonBytes, err := json.Marshal(students)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil

}

func exercise_1_1() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	defer file.Close()

	jsonBytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Ошибка при чтении файла", err)
		return
	}
	result, _ := modifyJSON(jsonBytes)
	fmt.Println(string(result))

}
