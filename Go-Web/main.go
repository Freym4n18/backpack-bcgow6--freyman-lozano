package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"io"
	"encoding/json"
	"os"
)

type Product struct {
	Name		string		`json:"name"`
    Price		float64		`json:"price"`
	Id			string		`json:"id"`
	Color		string		`json:"color"`
	Stock		int 		`json:"stock"`
	Code		int			`json:"code"`
	Published	bool		`json:"published"`
	CreateDate	time.Time	`json:"create_date"`
}

func check(err error) {
	if err!= nil {
        panic(err)
    }
}

func GetAll(c *gin.Context) {
	products := []Product{}
	jsonFile, err := os.Open("./Products.json")
	check(err)
	productByteArray, err := io.ReadAll(jsonFile)
	check(err)
    err = json.Unmarshal(productByteArray, &products)
	check(err)
	c.JSON(200,&products)
}

func main() {
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/products", GetAll)

	router.Run(":8080")
}
