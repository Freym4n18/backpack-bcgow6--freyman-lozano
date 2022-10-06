package handler

import (
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/pkg/web"
)
type Product struct {
	service products.Service
}

func NewProduct(service products.Service) *Product {
    return &Product{
		service: service,
	}
}

type ProductsPatchRequest struct {
	Name        string        `json:"name" binding:"required"`
    Price        float64        `json:"price" binding:"required"`
}


func CheckHeader(c *gin.Context) bool {
	err := godotenv.Load()
	if err!= nil {
		c.JSON(http.StatusInternalServerError, "Error loading environment: " + err.Error())
		return false
    }
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Unauthorized, missing authorization header"))
		return false
	}
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Unauthorized, invalid token"))
		return false
	}
	return true
}


func (p *Product) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
        if !CheckHeader(c) {
			return
		}
		var product products.Product
		err := c.BindJSON(&product)
		if err!= nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		err = p.service.AddProduct(product)

		if err!= nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}
