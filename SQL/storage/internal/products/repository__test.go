package products

import (
	"testing"

	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/SQL/storage/internal/domain"
	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/SQL/storage/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	//arrange
	product := domain.Product{
		Id: 4,
		Name: "PC",
		Type: "Technology",
		Count: 3,
		Price: 2000.0,
	}
	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	result,err := repo.GetByName(product.Name)
	assert.NoError(t,err)
	assert.Equal(t, product, result)
}

func TestGetAll(t *testing.T) {
	//arrange
	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	result,err := repo.GetAll()
	assert.NoError(t,err)
	assert.Greater(t, len(result), 0)
}

func TestUpdate(t *testing.T) {
	//arrange
	product := domain.Product{
		Name: "Laptop",
		Type: "Technology",
		Count: 3,
		Price: 2000.0,
	
	}
	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	result,err := repo.Store(product)
	assert.NoError(t,err)
	assert.Greater(t, int(result), 0)
	product.Id = result
	product.Name = "Update"
	product.Count = 4
	err = repo.Update(product, int(product.Id))
	assert.NoError(t, err)
	updated,err := repo.GetById(int(result))
	assert.NoError(t, err)
	assert.Equal(t, product, updated)
}

func TestDelete(t *testing.T) {
	//arrange
	product := domain.Product{
		Name: "Laptop",
		Type: "Technology",
		Count: 3,
		Price: 2000.0,
	
	}
	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	result,err := repo.Store(product)
	assert.NoError(t,err)
	assert.Greater(t, int(result), 0)

	err = repo.Delete(int(result))
	assert.NoError(t,err)
	product, err = repo.GetById(int(result))
	assert.Error(t, err)
	assert.Empty(t, product)
}

func TestStore(t *testing.T) {
	//arrange
	product := domain.Product{
		Name: "Laptop",
		Type: "Technology",
		Count: 3,
		Price: 2000.0,
	
	}
	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	result,err := repo.Store(product)
	assert.NoError(t,err)
	assert.Greater(t, int(result), 0)
}