package products

import (
	"encoding/json"
	"testing"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/Go-Web/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*Ejercicio 1 - Service/Repo/Db Update()

Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
	3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con
	mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test.
*/

func TestUpdateService(t *testing.T) {
	products := []*Product{
		{
			Id:     "1",
			Name: "Laptop",
			Stock:  12,
			Price: 100.0,
		},
		{
			Id:     "2",
			Name: "PC",
			Stock:  1,
			Price: 20,
		},
	}

	data, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	expProduct := Product{
		Id:     "1",
		Name: "Portatil",
		Stock:  12,
		Price: 100,
	}

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	product, err := service.UpdateProduct("1", "Portatil", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, expProduct, product)
	assert.Equal(t, dbMock.ReadInvoked, true)
}


func TestDelete(t *testing.T) {
	products := []*Product{
		{
			Id:     "1",
			Name: "Laptop",
			Stock:  12,
			Price: 100,
		},
		{
			Id:     "2",
			Name: "PC",
			Stock:  1,
			Price: 20,
		},
	}

	data, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	err := service.DeleteProduct("1")
	assert.Nil(t, err)
} 
