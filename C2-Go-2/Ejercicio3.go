package main

import (
	"fmt"
)

const (
	small = "small"
	medium = "medium"
    large = "large"
)

type ECommerce interface {
	Total() float64
	Add(product Producto)
}

type Store struct {
	Products []Producto
}

func (s *Store) Total() (total float64){
	for _,item := range s.Products{
		amount,err := item.calculateAmount()
		check(err)
		total += amount
	}
	return
}

func (s *Store) Add(produc Producto){
	s.Products = append(s.Products,produc)
}


type Producto interface{
	calculateAmount() (float64, error)
}


type Product struct {
	ProductType string
	Price float64
	Name string
}

func (p *Product) calculateAmount() (amount float64, err error) {

	switch p.ProductType {
    case small:
		amount = p.Price
	case medium:
		amount = p.Price * 1.03
	case large:
		amount = p.Price * 1.06 + 2500.0
	default:
		err = fmt.Errorf("Product type not found")
    }
	return 
}

func check(err error) {
	if err!= nil {
        panic(err)
    }
}

func NewProduct(productType string, price float64, name string) Producto {
	return &Product{
        ProductType: productType,
        Price:       price,
        Name:        name,
    }
}

func newEcommerce() ECommerce{
	return &Store{}
}

func main() {
	ecommerce := newEcommerce()
    ecommerce.Add(NewProduct("small", 3000, "Small"))
    ecommerce.Add(NewProduct("medium", 50000, "Medium"))
	fmt.Println(ecommerce.Total())
}
