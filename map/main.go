package main

import "fmt"

func main() {
	// users := map[string]int{
	// 	"vas": 10,
	// 	"pet": 20,
	// 	"max": 40,
	// }

	// users := make(map[string]int) // map нужно сразу инициализировать при создании
	// or
	var users map[string]int
	users = make(map[string]int)

	users["ser"] = 30
	// delete(users, "max")

	for key, value := range users {
		fmt.Println(key, value)
	}

	age, exist := users["max"]
	fmt.Println(age, exist)
}
