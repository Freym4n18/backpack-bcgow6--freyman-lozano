package products

import (
	"database/sql"

	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/SQL/storage/internal/domain"
)

const (
	GET_PRODUCT_BY_NAME = "SELECT id, name, type, count, price FROM products WHERE name=?;"
	SAVE_PRODUCT =  "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?);"

	GET_ALL_MOVIES = ""

	GET_MOVIE = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE id=?;"

	UPDATE_MOVIE = "UPDATE movies SET title=?, rating=?, awards=?, length=?, genre_id=? WHERE id=?;"

	DELETE_MOVIE = ""

	EXIST_MOVIE = "SELECT m.id FROM movies m WHERE m.id=?"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(p domain.Product) (int64, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_NAME, name)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Store(p domain.Product) (int64, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}