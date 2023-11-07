// Мы решили посчитать аналитику и посмотреть,
// а сколько же с нами учится каждый студент курса по Go -
// то есть найти кол-во дней, которое он (студент) провел в курсе
//  с момента поступления и до 1 октября 2023 года.

// Напишите функцию processJSON(jsonData []byte) error,
// которая должна принимать данные о студентах в формате JSON, разбирать их и выводить искомое число дней.

// Вывод должен быть в формате имяСтудента : количество дней

// Формат ввода
// [
//     {
//         "name": "Анна",
//         "admission_date": "2023-09-29T00:00:00Z"
//     },
//     {
//         "name": "Иван",
//         "admission_date": "2023-09-28T00:00:00Z"
//     }
// ]
// Формат вывода
// Анна: 2
// Иван: 3
// Примечания
// type Student struct {
//     Name         string    `json:"name"`
//     AdmissionDate time.Time `json:"admission_date"`
//     DaysOnCourse int       `json:"-"`
// }

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student4 struct {
	Name          string    `json:"name"`
	AdmissionDate time.Time `json:"admission_date"`
	DaysOnCourse  int       `json:"-"`
}

func processJSON(jsonData []byte) error {
	// create an array of structure
	var student4 []Student4

	// map json to the array of structure
	error := json.Unmarshal(jsonData, &student4)

	dateString := "2023-10-01T00:00:00Z"
	dateTime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	for i := 0; i < len(student4); i++ {
		// student4[i].DaysOnCourse = dateTime - student4[i].AdmissionDate
		student4[i].DaysOnCourse = int(dateTime.Sub(student4[i].AdmissionDate).Hours() / 24)
		fmt.Printf("%s: %v \n", student4[i].Name, student4[i].DaysOnCourse)
	}

	return error

}

func Exercise_1_4() {
	inputData := []byte(`[
		{
			"name": "Анна",
			"admission_date": "2023-09-29T00:00:00Z"
		},
		{
			"name": "Иван",
			"admission_date": "2023-09-28T00:00:00Z"
		}
	]`)

	processJSON(inputData)
}
