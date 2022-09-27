package main

import (
	"fmt"
	"time"
)

type Student struct {
	FirstName	string
	LastName	string
	DNI			string
	Date		string
}

func(s Student) Detalle() {
	fmt.Println("First Name:",s.FirstName)
	fmt.Println("Last Name:",s.LastName)
	fmt.Println("DNI:",s.DNI)
	fmt.Println("Date: ",s.Date)
}

func main() {
	currentTime := time.Now()
	fmt.Println(TypeOf(currentTime));
}
