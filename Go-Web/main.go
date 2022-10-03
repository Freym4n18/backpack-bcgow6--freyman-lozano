package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"io"
	"encoding/json"
	"os"
	"net/http"
	"strconv"
	"fmt"
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

	//lets start with the filters
	name := c.Request.URL.Query().Get("name")
	color := c.Request.URL.Query().Get("color")
	stock := c.Request.URL.Query().Get("stock")
	code := c.Request.URL.Query().Get("code")
	published := c.Request.URL.Query().Get("published")
	afterDate := c.Request.URL.Query().Get("after_date")
	price := c.Request.URL.Query().Get("price")
	

	fmt.Println(name)
	productsFiltered := []Product{}

	for _, product := range products {
		addproduct := true
		if name!= "" {
            if product.Name!= name {
				addproduct = false
			}
		}
		if color!= "" {
            if product.Color!= color {
				addproduct = false
			}
		}
		if stock!= "" {
			stockInt, err := strconv.Atoi(stock)
			check(err)
            if product.Stock!= stockInt {
				addproduct = false
            }
		}
		if code!= "" {
			codeInt, err := strconv.Atoi(code)
			check(err)
            if product.Code!= codeInt {
				addproduct = false
			}
		}
		if published!= "" {
			publishedBool, err := strconv.ParseBool(published)
            check(err)
            if product.Published!= publishedBool {
                addproduct = false
            }
		}
		if afterDate!= "" {
			createDateConv, err := time.Parse("2006-01-02", afterDate)
            check(err)
            if product.CreateDate.Before(createDateConv) {
                addproduct = false
            }
		}
		if price!= "" {
			priceFloat, err := strconv.ParseFloat(price, 64)
            check(err)
            if product.Price!= priceFloat {
                addproduct = false
            }
		}
		if addproduct {
            productsFiltered = append(productsFiltered, product)
		}
	}

	c.JSON(http.StatusAccepted,&productsFiltered)
}

func GetOne(c *gin.Context) {
	id := c.Param("id")
    products := []Product{}
    jsonFile, err := os.Open("./Products.json")
    check(err)
    productByteArray, err := io.ReadAll(jsonFile)
    check(err)
	err = json.Unmarshal(productByteArray, &products)
    check(err)
	for _, product := range products {
        if product.Id == id {
			c.JSON(200,&product)
            return
		}
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Product ID not found",
	})
}

func main() {
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/products", GetAll)

	router.GET("/products/:id", GetOne)

	router.Run(":8080")
}
