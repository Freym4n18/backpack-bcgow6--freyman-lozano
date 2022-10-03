package main

import (
	"errors"
	"fmt"
)

func main() {
	mostrarPromedio(getPromedio(1,2,3))
}

func mostrarPromedio(promedio float64, err error) {
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(promedio)
	}
}

func getPromedio(nums ...float64) (promedio float64, err error){
	for _, value := range nums {
		if value < 0.0 {
			return -1, errors.New("No puede haber calificaciones negativas")
		}	
		promedio += value
	}
	promedio /= float64(len(nums))
	return
}