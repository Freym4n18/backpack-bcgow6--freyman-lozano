package products

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{
	Data []Product
}

func (fs *StubStore) Read(data interface{}) error {
	a := data.(*[]Product)
	*a = fs.Data
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	expected := []Product{
		{
			Id:     "1",
			Name: "PC",
			Stock:  10,
			Price: 10.0,
		},
		{
			Id:     "2",
			Name: "Laptop",
			Stock:  23,
			Price: 10.0,
		},
	}
	
	stub := &StubStore{}
	stub.Data = expected
	repo := NewRepository(stub)

	a, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, a)
}

type MockStore struct {
	ReadInvoked bool
	Data        []Product
}

func (fs *MockStore) Read(data interface{}) error {
	fs.ReadInvoked = true
	a := data.(*[]Product)
	*a = fs.Data
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	return nil
}

func TestUpdate(t *testing.T) {
	id, newNombre, newPrice := "1", "Update After", 10.0
	products := []Product{{Id: "1", Name: "Update Before", Stock: 1, Price: 12}}

	mock := MockStore{Data: products}

	r := NewRepository(&mock)
	productUpdated, err := r.Update(id,newNombre, newPrice)
	assert.Nil(t, err)

	assert.Equal(t, id, productUpdated.Id)
	assert.Equal(t, newNombre, productUpdated.Name)
	assert.Equal(t, newPrice, productUpdated.Price)
	assert.True(t, true, mock.ReadInvoked)
}
