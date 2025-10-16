package model

type Product struct {
	Id    uint32
	Name  string
	Price float32
}

type ProductRepository interface {
	FindByID(id int) (*Product, error)
	FindAll() ([]*Product, error)
	Create(product *Product) error
}
