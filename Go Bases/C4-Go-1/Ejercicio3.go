package main

import (
	"fmt"
)

type myCustomError struct {
	status int
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.status, e.msg)
}

func CheckSalary(salary int) {
	customError := fmt.Errorf("Salary must be greater than 150000, got %d", salary)
	if(salary < 150000) {
		fmt.Println(customError.Error())
		return
	}
	fmt.Println("Debe Pagar impuesto")
}

func main() {
	salary := 140000
	CheckSalary(salary)
}
