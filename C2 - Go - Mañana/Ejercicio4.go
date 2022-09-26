package main

import (
	"fmt"
	"math"
	"errors"
)

func getAverage(nums []float64) float64 {
	average := 0.0
	for _,value := range nums {
		average += value
	}
	average /= float64(len(nums))
	return average
}

func getMininum(nums []float64) float64 {
	minimun := 1000000000.0 //a big number
	for _,value := range nums {
		minimun = math.Min(value,minimun)
	}
	return minimun
}

func getMaximun(nums []float64) float64 {
	maximun := 0.0
	for _,value := range nums {
		maximun = math.Max(maximun,value)
	}
	return maximun
}

var funciones = map[string]func(nums []float64)float64{"Average":getAverage,"Minimum":getMininum,"Maximum":getMaximun}

func calcularFuncion(funcion string, nums ...float64) (result float64, err error) {
	var ok bool
	_, ok = funciones[funcion]
	if !ok {
		err = errors.New("No existe la funcion especificada")
		return
	}
	operacion := funciones[funcion]
	result = operacion(nums)
	return
}

func mostrarResultado(result float64, err error) {
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func main() {
	result, err := calcularFuncion("Average",2,3,4,5)
	mostrarResultado(result,err)

	result, err = calcularFuncion("sum",2,3,4,5)
	mostrarResultado(result,err)

	result, err = calcularFuncion("Minimum",2,3,4,5)
	mostrarResultado(result,err)

}
