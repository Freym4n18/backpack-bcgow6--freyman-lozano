package main

import (
	"fmt"
)

var Salario = map[string]int{"A":3000, "B":1500, "C":1000}
var PorcentajeAdicional = map[string]float64{"A":0.5, "B":0.2, "C":0.0}

func main() {
	fmt.Println(calcularSalario(20,"A"));
}


func calcularSalario(minutosTrabajados int, categoria string) (salarioAdicional float64) {
	salarioBase := minutosTrabajados * Salario[categoria]
	salarioAdicional = float64(salarioBase) * (1.0 + PorcentajeAdicional[categoria])
	return
}