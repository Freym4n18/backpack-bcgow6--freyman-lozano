package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	//consultamos a Benjamin
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	//calculamos la cantidad de empleados mayores a 21 años
	n := 0
	for _, value := range employees {
		if value >= 21 {
			n++;
		}
	}
	fmt.Println("Hay",n,"Empleados mayores a 21 años.");

	//agregar a Federico con 25 años
	employees["Federico"] = 25

	//eliminamos a Pedro
	delete(employees, "Pedro")

	fmt.Println("La lista actualizada de empleados es:")
	fmt.Println(employees)


}
