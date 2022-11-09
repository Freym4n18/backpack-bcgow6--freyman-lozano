package products

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Freym4n18/backpack-bcgow6--freyman-lozano/SQL/storage/internal/domain"
)

const (
	GET_PRODUCT_BY_NAME = "SELECT id, name, type, count, price FROM products WHERE name=?;"

	GET_PRODUCT_BY_ID = "SELECT id, name, type, count, price FROM products WHERE id=?;"
	
	SAVE_PRODUCT =  "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?);"

	GET_ALL_PRODUCTS = "SELECT id, name, type, count, price FROM products"

	UPDATE_PRODUCT = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"

	DELETE_PRODUCT = "DELETE FROM products WHERE id=?"

)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(p domain.Product) (int64, error)
	Update(p domain.Product, id int) (error)
	Delete(id int) (error)
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
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

func (r *repository) GetById(id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_ID, id)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(id int) (error) {
	stmt, err := r.db.Prepare(DELETE_PRODUCT)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}


func (r *repository) Update(p domain.Product, id int) error {
	stm, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
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