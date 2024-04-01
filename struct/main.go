package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex)
}

func NewUser(name, sex string, age int) User {
	return User{
		name: name,
		age:  age,
		sex:  sex,
	}
}

func main() {

	user := User{"sergey", 33, "male"}
	user2 := NewUser("max", "male", 25)

	fmt.Printf("%+v\n", user)
	fmt.Println(user2.name, user2.age)
	user2.printUserInfo()
}
