package products

import (
	"strconv"
	"time"
	"errors"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/pkg/store"
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

type Repository interface {
	GetAll()(products []Product, err error)
	GetById(id string) (product Product, err error)
	Update(id string, name string, price float64) (err error)
	Delete(id string) (err error)
	Replace(id string, product Product) (err error)
	FindNextId() (id string, err error)
	AddProduct(producto Product) (err error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
    return &repository{
		db: db,
	}
}


func (r *repository) GetAll() (products []Product, err error) {
	products = []Product{}
	err = r.db.Read(&products)
	if err!= nil {
		return products, err
	}
	return
}

func (r *repository) GetById(id string) (product Product, err error) {
	products := []Product{}
	err = r.db.Read(&products)
	if err!= nil {
        return
    }
	idx, err := findProduct(products, id)
	if err!= nil {
        return Product{}, err
    }
	product = products[idx]
    return
}

func (r *repository) Update(id string, name string, price float64) (err error) {
	products := []Product{}
	r.db.Read(&products)
	idx, err := findProduct(products,id)
	if err!= nil {
        return
    }
	products[idx].Name = name
	products[idx].Price = price
    r.db.Write(&products)
	return
}

func (r *repository) Delete(id string) (err error) {
	products := []Product{}
	r.db.Read(&products)
    idx, err := findProduct(products, id)
	if err!= nil {
		return err
	}
	products = append(products[:idx], products[idx+1:]...)
	r.db.Write(products)
    return nil
}

func (r *repository) Replace(id string, product Product) (err error) {
	products := []Product{}
	r.db.Read(&products)
	idx, err := findProduct(products, id)
	product.Id = id
	products[idx] = product
	r.db.Write(products)
	return
}

func (r *repository) FindNextId() (id string, err error) {
	products := []Product{}
	r.db.Read(&products)
	
	idint := findId(products)
	if err!= nil {
        return "", err
    }
	id = strconv.Itoa(idint)
	return
}

func findId(products []Product) int {
	maxId := 0
	for _, product := range products {
		id, err := strconv.Atoi(product.Id)
		if err!= nil {
			panic(err)
		}
        if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}


func (r *repository) AddProduct(producto Product) (err error) {
	products := []Product{}
	r.db.Read(&products)
	products = append(products, producto)
	err = r.db.Write(products)
    return err
}

func findProduct(products []Product, id string) (int, error) {
	for idx, product := range products {
        if product.Id == id {
            return idx, nil
        }
    }
	return 0, errors.New("Product not found with id: " + id)
 }



