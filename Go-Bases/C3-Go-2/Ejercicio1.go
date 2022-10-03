package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
	Email string
	Password string
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func(u *User) SetEmail(email string) {
	u.Email = email
}

func(u *User) SetPassword(password string) {
	u.Password = password
}

func main() {
	user := User {
		Name: "Freyman",
		Age: 23,
		Email: "freiman@gmail.com",
		Password: "superpassword3000",
	}

	fmt.Printf("%+v\n",user);

	user.SetAge(14)
	user.SetName("Yohani")

	fmt.Printf("%+v\n",user)
}
