package main

import (
	"fmt"
)

func main() {
	msg := "hello"
	fmt.Println(msg)
	printMsg(&msg) // & - мы ссылаемся на область памяти
	fmt.Println(msg)

	users := make([]string, 10)
	fmt.Println(users)
}

func printMsg(msg *string) {
	*msg += " cfcf" // * - мы берем значение из области памяти
}
