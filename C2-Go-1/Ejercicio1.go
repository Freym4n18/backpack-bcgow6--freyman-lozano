package main

import (
	"fmt"
)

const (
	primerImpuesto = 0.17 
	adicional = 0.10
)

func main() {
	fmt.Println("El impuesto para un salario de 20000 es" ,getImpuesto(20000.0))
	fmt.Println("El impuesto para un salario de 60000 es" ,getImpuesto(60000.0))
	fmt.Println("El impuesto para un salario de 160000 es" ,getImpuesto(160000.0))
}

func getImpuesto(salario float64 ) float64 {
	impuestoAcumulado := 0.0
	if salario > 50000.0 {
		impuestoAcumulado += primerImpuesto
	}
	if salario > 150000.0 {
		impuestoAcumulado += adicional
	}
	return salario * impuestoAcumulado
}