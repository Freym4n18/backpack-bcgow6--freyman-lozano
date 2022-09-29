package main

import (
	"fmt"
)


type Product struct {
	Name string
	Price float64
	Amount int
}

type Service struct{
	Name string
	Price float64
	WorkedMinutes int
}

type Maintenance struct{
	Name string
	Price float64
}

func SumProducts(array []Product, ch chan float64) {
	total := 0.0
	for _,product := range array{
		total += product.Price * float64(product.Amount)
	}
	ch <- total
}

func SumServices(array []Service, ch chan float64) {
	total := 0.0
	for _,service := range array {
		ciclesWorked := (service.WorkedMinutes + 29)/30
		total += float64(ciclesWorked) * service.Price
	}
	ch <- total
}

func sumMaintenances(array []Maintenance, ch chan float64) {
	total := 0.0
	for _,maintenance := range array{
		total += maintenance.Price
	}
	ch <- total
}

func main() {
	canal := make(chan float64)
	s1 := Service{
		Name: "Service1",
		Price: 2.5,
		WorkedMinutes: 120,
	}

	s2 := Service{
		Name: "Service2",
		Price: 10.5,
		WorkedMinutes: 29,
	}

	m1 := Maintenance{
		Name: "Maintenance1",
		Price: 50.0,
	}


	m2 := Maintenance{
		Name: "Maintenance2",
		Price: 150.0,
	}

	p1 := Product{
		Name: "Producto1",
		Price: 12.4,
		Amount: 3,
	}

	p2 := Product{
		Name: "Producto2",
		Price: 30.0,
		Amount: 5,
	}

	services := []Service{s1,s2}
	products := []Product{p1,p2}
	maintenances := []Maintenance{m1,m2}

	totalAmount := 0.0

	go SumProducts(products,canal)
	go SumServices(services,canal)
	go sumMaintenances(maintenances,canal)

	result1 := <- canal
	result2 := <- canal
	result3 := <- canal

	totalAmount = result1 + result2 + result3
	fmt.Println(totalAmount);
}
