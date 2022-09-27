package main

import (
	"fmt"
	"time"
)

type Student struct {
	FirstName	string
	LastName	string
	DNI			string
	Date		time.Time
}

func(s Student) Detalle() {
	fmt.Println("First Name:",s.FirstName)
	fmt.Println("Last Name:",s.LastName)
	fmt.Println("DNI:",s.DNI)
	fmt.Println("Date: ",s.Date)
}

func main() {
	s := Student{
		FirstName: "Freyman"
		LastName: "Lozano"
		DNI: "23432FD"
		Date: time.Now()
	}
	fmt.Println(s);
}
