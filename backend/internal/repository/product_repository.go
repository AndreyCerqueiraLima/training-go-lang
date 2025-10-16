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

type ProductSQLRepository struct {
	sql.DB
}

func NewProductRepository() ProductRepository {
	return &ProductSQLRepository{}
}

func (r *ProductSQLRepository) FindById(id int) (*model.Product, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/training_test_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var p model.Product
	db.QueryRow("SELECT id,name,price from products where id = ?", id).Scan(&p.Id, &p.Name, &p.Price)

	return &p, errors.New("não encontrado")
}

func (r *ProductSQLRepository) FindAll() ([]*model.Product, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/training_test_go")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var products []*model.Product

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *ProductSQLRepository) Create(p model.Product) error {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/training_test_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err2 := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)

	if err2 != nil {
		return err2
	}

	return nil
}
