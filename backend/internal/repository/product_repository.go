package repository

import (
	"bjj-system/internal/model"
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ProductRepository interface {
	FindById(id uint32) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Create(product *model.Product) error
	Delete(id uint32) error
}

// implements ProductRepository
type ProductSQLRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductSQLRepository{db: db}
}

func (r *ProductSQLRepository) FindById(id uint32) (*model.Product, error) {
	var p model.Product
	r.db.QueryRow("SELECT id,name,price from products where id = ?", id).Scan(&p.Id, &p.Name, &p.Price)

	if p.Name == "" {
		return nil, errors.New("não encontrado")
	}

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

	return products, nil
}

func (r *ProductSQLRepository) Create(p *model.Product) error {
	result, err := r.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = uint32(id)

	return nil
}

func (r *ProductSQLRepository) Delete(id uint32) error {
	result, err := r.db.Exec("DELETE products WHERE id = ?", id)

	if err != nil {
		return err
	}
	rowsAfectNumb, _ := result.RowsAffected()

	if rowsAfectNumb > 0 {
		return nil
	}

	return errors.New("nenhum registro deletado")
}
