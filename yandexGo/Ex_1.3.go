// Равнение на флаг

// На линейке ученикам нужно сгруппироваться по классам.

// Проведём линейку для базы данных.

// Напишите функцию splitJSONByClass(jsonData []byte) (map[string][]byte, error), которая принимает данные в формате JSON и возвращает мапу,
// в которой ключи — классы, а значения — данные в формате JSON, которые к этому классу относятся.

// Примечания
// Например: Входные данные

// [
//     {
//         "name": "Oleg",
//         "class": "9B"
//     },
//     {
//         "name": "Ivan",
//         "class": "9A"
//     },
//     {
//         "name": "Maria",
//         "class": "9B"
//     },
//     {
//         "name": "John",
//         "class": "9A"
//     }
// ]
// Выходные данные должны быть в виде map:

// map[string][]byte{
//         "9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
//         "9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]`),
// }

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var students []Student

	// JSON --> Go
	if err := json.Unmarshal(jsonData, &students); err != nil {
		return nil, err
	}

	result := make(map[string][]byte)
	for _, student := range students {
		// Go --> JSONbytes
		bytes, err := json.Marshal(student)
		if err != nil {
			return nil, err
		}

		classData := result[student.Class]
		if classData == nil {
			result[student.Class] = bytes
		} else {
			result[student.Class] = append(classData, bytes...)
		}
	}

	return result, nil
}
func Exercise_1_3() {

	inputData := []byte(`[
		    {
		        "name": "Oleg",
		        "class": "9B"
		    },
		    {
		        "name": "Ivan",
		        "class": "9A"
		    },
		    {
		        "name": "Maria",
		        "class": "9B"
		    },
		    {
		        "name": "John",
		        "class": "9A"
		    }
		]`)

	splittedJSONbyClass, _ := splitJSONByClass(inputData)

	for class, data := range splittedJSONbyClass {
		fmt.Println(class, string(data))
	}
}
