package main

import (
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/cmd/server/handler"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/internal/products"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	db := store.NewFileStore("./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.GET("",p.GetAll())
	pr.GET("/:id",p.GetOne())
	pr.POST("", p.AddProduct())
	pr.PATCH("/:id",p.UpdateProduct())
	pr.DELETE("/:id",p.DeleteProduct())
	pr.PUT("/:id",p.ReplaceProduct())
	r.Run()
}