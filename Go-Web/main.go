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
	"errors"
)

type Product struct {
	Name		string		`json:"name" binding:"required"`
    Price		float64		`json:"price" binding:"required"`
	Id			string		`json:"id"`
	Color		string		`json:"color" binding:"required"`
	Stock		int 		`json:"stock" binding:"required"`
	Code		int			`json:"code" binding:"required"`
	Published	bool		`json:"published" binding:"required"`
	CreateDate	time.Time	`json:"create_date" binding:"required"`
}

type ProductsPatchRequest struct {
	Name        string        `json:"name" binding:"required"`
    Price        float64        `json:"price" binding:"required"`
}

func check(err error) {
	if err!= nil {
    	fmt.Println(err.Error())
    }
}

func CheckHeader(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
            "code": 401,
            "msg":  "Unauthorized, missing authorization header",
        })
		return false
	}
	if token != "abc123" {
		c.JSON(http.StatusUnauthorized, gin.H{
            "code": 401,
            "msg":  "Unauthorized, invalid authorization header",
        })
		return false
	}
	return true
}

func findId(products []Product) int {
	maxId := 0
	for _, product := range products {
		id, err := strconv.Atoi(product.Id)
		check(err)
        if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

func ReadJson() []Product{
	products := []Product{}
    jsonFile, err := os.Open("./Products.json")
    check(err)
    productByteArray, err := io.ReadAll(jsonFile)
    check(err)
	err = json.Unmarshal(productByteArray, &products)
    check(err)
	return products
}

func WriteJson(products []Product) {
	jsonFile, err := os.Create("./Products.json")
    check(err)
    defer jsonFile.Close()
    jsonData, err := json.MarshalIndent(products, "", "\t")
    check(err)
	_, err = jsonFile.Write(jsonData)
    check(err)
	jsonFile.Sync()
}

func findProduct(products []Product, id string) (int, error) {
	for idx, product := range products {
        if product.Id == id {
            return idx, nil
        }
    }
	return 0, errors.New("Product not found with id: " + id)
 }

func GetAll(c *gin.Context) {
	if !CheckHeader(c) {
		return
	}
	products := ReadJson()

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
	//validate the parameters
	if !CheckHeader(c) {
		return
	}
	id := c.Param("id")

	//Get the product with the given id

    products := ReadJson()

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



func Create(c *gin.Context) {
	//validate the parameters
	if !CheckHeader(c) {
		return
	}
	var product Product
	err := c.BindJSON(&product)
    if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//create product
	products := ReadJson()
	id := findId(products)
	product.Id = strconv.Itoa(id)
	products = append(products,product)
	
	WriteJson(products)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added",
	})
}

 func UpdateOne(c *gin.Context) {
	//validate the parameters
	if!CheckHeader(c) {
        return
    }

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": "Product ID not found",
        })
		return
	}

	var product Product
    err := c.BindJSON(&product)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
		return
	}

	//update product
	products := ReadJson()

    idx, err := findProduct(products, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
		return
	}
	product.Id = id
    products[idx] = product

	WriteJson(products)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added",
	})
}

func DeleteOne(c *gin.Context) {
	//validate the parameters
    if!CheckHeader(c) {
        return
    }

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": "Product ID not found",
        })
		return
	}

	products := ReadJson()
    idx, err := findProduct(products, id)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
		return
	}
	products = append(products[:idx], products[idx+1:]...)

	WriteJson(products)

	c.JSON(http.StatusOK, &products[idx])
}

func PatchOne(c *gin.Context) {
	//validate the parameters
    if!CheckHeader(c) {
        return
    }

	id := c.Param("id")
    if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product ID not found",
		})
		return
	}

	var req ProductsPatchRequest
	err := c.BindJSON(&req)
    if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//patch product
	products := ReadJson()
	idx, err := findProduct(products, id)
    if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	products[idx].Name = req.Name
	
	products[idx].Price = req.Price
	check(err)

	WriteJson(products)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product patched",
	})
}

func main() {
	router := gin.Default()

	router.GET("/products", GetAll)

	router.GET("/products/:id", GetOne)

	router.PUT("/products/:id", UpdateOne)

	router.DELETE("/products/:id", DeleteOne)

	router.POST("/products", Create)

	router.PATCH("/products/:id", PatchOne)

	router.Run(":8080")
}
