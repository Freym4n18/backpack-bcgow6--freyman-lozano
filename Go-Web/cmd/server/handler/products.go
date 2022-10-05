package handler

import (
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/intenal/products"
)
type Product struct {
	service products.Service
}

func NewProduct(service products.Service) *Product {
    return &Product{
		service: service,
	}
}
