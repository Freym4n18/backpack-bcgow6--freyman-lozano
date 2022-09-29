package main

import (
	"fmt"
)

func main() {
	edad := 22
	antiguedad := 1
	sueldo := 9999
	switch {
	case edad < 22  || antiguedad < 1:
		fmt.Println("No se te puede otorgar el prestamo debido a que no cumples los requisitos necesarios.")
	case sueldo <= 100000:
		fmt.Println("Tu prestamo sera otorgado sin intereses.")
	case sueldo >= 100000:
		fmt.Println("Tu prestamo sera otorgado con intereses debido a tu alta capacidad salarial");
	}
}
