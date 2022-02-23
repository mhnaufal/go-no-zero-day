package day_25

import (
	"encoding/json"
	"fmt"
)

// create a struct with structure match the JSON format
type User struct {
	FullName string `json:"Name"` //tag json:"Name". Used to rename the json property w/ struct property
	Age      int
}

func Json() {
	var jsonString = `{"Name": "Hyper", "Age": 33}`
	var jsonData = []byte(jsonString) //convert jsonString into array of bytes

	fmt.Println("Json data: ", jsonString)
	fmt.Println("Json data: ", jsonData)

	var data User

	// Decode JSON to struct or object
	// Unmarshall only accept data in byte format
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name	: ", data.FullName)
	fmt.Println("Age	: ", data.Age)

	// Decode to interface
	var data1 map[string]interface{}
	var err1 = json.Unmarshal(jsonData, &data1)

	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	fmt.Println("Name 2	: ", data1["Name"]) //watch out we called it using key "Name"
	fmt.Println("Age 2	: ", data1["Age"])
}
