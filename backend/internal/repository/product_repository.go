package repository

import (
	"bjj-system/internal/model"
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ProductRepository interface {
	FindById(id int) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Create(product model.Product) error
}

// implements ProductRepository
type ProductSQLRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductSQLRepository{db: db}
}

func (r *ProductSQLRepository) FindById(id int) (*model.Product, error) {
	var p model.Product
	r.db.QueryRow("SELECT id,name,price from products where id = ?", id).Scan(&p.Id, &p.Name, &p.Price)

	if p.Name == "" {
		return nil, errors.New("não encontrado")
	}

	defer r.db.Close()

	return &p, nil
}

func (r *ProductSQLRepository) FindAll() ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}

	var products []*model.Product

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, &p)
	}
	defer r.db.Close()

	return products, nil
}

func (r *ProductSQLRepository) Create(p model.Product) error {
	_, err2 := r.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)

	if err2 != nil {
		return err2
	}
	defer r.db.Close()

	return nil
}
