package main

import (
	"fmt"
)

func main() {
	mes := 7

	//forma 1
	var meses = map[int]string{1:"Enero", 2:"Febrero", 3:"Marzo", 4:"Abril" , 5:"Mayo", 6:"Junio", 7:"Julio", 8:"Agosto", 9:"Septiembre", 10:"Octubre", 11:"Noviembre", 12:"Diciembre"}
	fmt.Println(meses[mes])

	//forma 2
	meses1 := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	fmt.Println(meses1[mes-1])

}
