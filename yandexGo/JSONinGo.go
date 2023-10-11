package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ----------------------------
// JSON --> Golang: Unmarshal()
//
// Golang --> JSON: Marshal()
// ----------------------------

// Структура Person с полями JSON
type Tourist struct {
	Name  string `json:"name"` // Имя в формате JSON
	Email string `json:"email"`
}

// Функция для открытия файла
func OpenFile() {
	// Открываем файл data.json.
	// Если будет ошибка открытия, то в переменную err что-то запишется.
	// Соответственно, если err будет не nil (null, None),
	// то туда что-то запишется.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	// defer will move the execution of the statement
	// to the very end inside a function.
	defer file.Close()

	// Reading the file
	byteValue, _ := io.ReadAll(file)

	// create a future instance of Person
	var traveller Tourist

	// Unmarshal JSON --> Person structure
	err = json.Unmarshal(byteValue, &traveller)

	if err != nil {
		fmt.Print("Unmarshaling error:", err)
	}

	fmt.Println("Name", traveller.Name)
	fmt.Println("Email", traveller.Email)

	// Marshal person to JSON
	jsonBytes, err := json.Marshal(traveller)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

}

func JsonEncoder() {
	var traveller Tourist

	// variable buffer to store JSON-data
	var buffer bytes.Buffer

	// encoder writes JSON to buffer
	encoder := json.NewEncoder(&buffer)

	err := encoder.Encode(traveller)
	if err != nil {
		fmt.Println("JSON data writing error", err)
		return
	}

	// print result
	fmt.Println("JSON data about the travellers:")
	fmt.Println(buffer.String())
}
