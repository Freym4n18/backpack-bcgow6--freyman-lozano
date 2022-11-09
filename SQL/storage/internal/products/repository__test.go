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