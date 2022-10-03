package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Product struct {
	Id string
	Precio float64
	Cantidad int
	PrecioTotal float64
}


func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("./Products.csv")
	check(err)
	fmt.Println(string(data))

	csvlinesList := strings.Split(string(data),"\n")

	fmt.Printf("%20s%10s%10.2s%10.2s\n","Id","Cantidad","Precio","PrecioTotal")
	for _,value := range csvlinesList{

		csvLine := strings.Split(value,",")

		if len(csvLine) != 3 {
			continue
		}

		id := csvLine[0]
		precio,_ := strconv.ParseFloat(csvLine[1],64)
		cantidad,_ := strconv.Atoi(csvLine[2])

		p := Product{
			Id: id,
			Precio: precio,
			Cantidad: cantidad,
			PrecioTotal: precio * float64(cantidad),
		}
		
		fmt.Printf("%20s%10d%10.2f%10.2f\n",p.Id,p.Cantidad,p.Precio,p.PrecioTotal)

	}
}
