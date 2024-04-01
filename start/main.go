package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func main() {

	hello := ""
	hello = "hellllllo"
	fmt.Println(hello)

	var msg string = "msg"
	fmt.Println(msg)

	const msg_2 = "msg_2"
	fmt.Println(reflect.TypeOf(msg_2))

	// однострочный комментарий
	/*
		много
		строчный
		комментарий
	*/

	// if else
	message, err := enterTheClub(19)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// switch
	day, err := prediction("m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day)

	//...int
	min_num := findMin(10, 40, 50, 5)
	fmt.Println(min_num)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "go", nil
	} else if age >= 45 {
		return "go, but...", nil
	}

	return "wait", errors.New("you are too young")
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "m":
		return "monday", nil
	case "t":
		return "tuesday", nil
	case "w":
		return "wednesday", nil
	default:
		return "week", errors.New("not valid")
	}
}

func findMin(num ...int) int {
	if len(num) == 0 {
		return 0
	}
	min := num[0]
	for _, i := range num {
		if i < min {
			min = i
		}
	}
	return min
}
