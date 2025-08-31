package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

const FILEPATH = "./hard/json-serialization/users.json"

func main() {
	users := []User{
		{Name: "Dmytro", Age: 18},
		{Name: "Andriy", Age: 21},
		{Name: "Yaroslav", Age: 23},
	}

	file, err := os.Create(FILEPATH)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(users); err != nil {
		panic(err)
	}

	fmt.Println("Data saved successfully in file users.json")

	file2, err := os.Open(FILEPATH)

	if err != nil {
		panic(err)
	}

	defer file2.Close()

	var loadedUsers []User

	decoder := json.NewDecoder(file2)

	if err := decoder.Decode(&loadedUsers); err != nil {
		panic(err)
	}

	fmt.Println("Data read successfully from:")
	for _, u := range loadedUsers {
		fmt.Printf("Name: %s, Age: %d\n", u.Name, u.Age)
	}
}
