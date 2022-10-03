package main

import (
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	Id string
	Precio float64
	Cantidad int
}

func (p Product) csvFormat() string {
	precioString := strconv.FormatFloat(p.Precio, 'E', -1, 64)
	cantidadString := strconv.Itoa(p.Cantidad)
	return p.Id +  "," + precioString + "," + cantidadString
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	p1 := Product{
		Id: "ZL Mouse 1234",
		Precio: 12.5,
		Cantidad: 5,
	}
	p2 := Product{
		Id: "ZL Teclado 123",
		Precio: 20,
		Cantidad: 5,
	}

	p3 := Product{
		Id: "RedDragon Silla 500P",
		Precio: 99.9,
		Cantidad: 3,
	}

	list := []Product{p1,p2,p3}

	f, err := os.Create("./Products.csv")
	check(err)
	defer f.Close()
	for _,value := range list {
		fmt.Printf("Writing produc with id: %s\n",value.Id)
		_, err := f.WriteString(value.csvFormat() + "\n")
		check(err)
		f.Sync()
	}
	fmt.Println("done.")
}
