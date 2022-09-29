package main

import (
	"fmt"
)

type Product struct {
	Name string
	Price float64
	Amount int
}


func New(name string, price float64) Product{
	p := Product{
		Name: name,
		Price: price,
	}
	return p
}

type User struct {
	FirstName string
	LastName string
	Email string
	Products []Product
}

func(u *User) AddProduct(p Product, amount int) {
	p.Amount = amount
	u.Products = append(u.Products,p)
}

func (u *User) DeleteProducts() {
	u.Products = []Product{}
	
}

func main() {

	user := User{
		FirstName: "Freyman",
		LastName: "Lozano",
		Email: "freiman@gmail.com",
	}

	p := New("teclado",2900)

	fmt.Printf("%+v\n",user)

	user.AddProduct(p,2)

	fmt.Printf("%+v\n",user);

	user.DeleteProducts()

	fmt.Printf("%+v\n",user)


}
