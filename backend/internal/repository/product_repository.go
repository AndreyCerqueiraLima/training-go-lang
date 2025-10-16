package repository

import (
	"bjj-system/internal/model"
	"database/sql"

	_ "modernc.org/sqlite"
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
	product := new(model.Product)
	return product, nil
}

func (r *ProductSQLRepository) FindAll() ([]*model.Product, error) {
	products := []*model.Product{
		{Name: "Camiseta", Price: 49.90},
		{Name: "Calça", Price: 99.90},
		{Name: "Tênis", Price: 199.90},
	}
	return products, nil
}

func (r *ProductSQLRepository) Create(product model.Product) error {

}
