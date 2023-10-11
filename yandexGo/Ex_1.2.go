package main

import (
	"encoding/json"
	"fmt"
)

type Pupil struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var pupil []Pupil
	var s []Pupil

	for i := 0; i < len(jsonDataList); i++ {
		json.Unmarshal(jsonDataList[i], &pupil)
		s = append(s, pupil...)
	}

	res, err := json.Marshal(s)
	return res, err
}

func Exercise_1_2() {
	first := []byte(`[
		{
			"name": "Oleg",
			"class": "9B"
		},
		{
			"name": "Ivan",
			"class": "9A"
		}
	]`)

	second := []byte(`[
		{
			"name": "Maria",
			"class": "9B"
		},
		{
			"name": "John",
			"class": "9A"
		}
	]`)

	res, _ := mergeJSONData(first, second)
	fmt.Println(string(res))

}
