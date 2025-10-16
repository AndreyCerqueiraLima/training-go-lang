package service

import (
	"bjj-system/internal/model"
	"bjj-system/internal/model/dto"
	"bjj-system/internal/repository"
)

type ProductService struct {
	p *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{p: repo}
}

func (s ProductService) CreateProduct(product dto.ProductIn) dto.ProductOut {
	//repository := new(repository.ProductSQLRepository)
	domainProduct := model.Product{}
	domainProduct.Name = product.Name
	domainProduct.Price = product.Price
	//repository.Create(domainProduct)
	productOut := dto.ProductOut{}
	productOut.Name = domainProduct.Name
	productOut.Price = domainProduct.Price
	return productOut
}

func (s ProductService) GetProducts() []dto.ProductOut {
	result := []dto.ProductOut{}
	repository := repository.ProductSQLRepository{}
	products, err := repository.FindAll()
	if err != nil {
		return result
	}

	for _, product := range products {
		result = append(result, dto.ProductOut{Name: product.Name, Price: product.Price})
	}

	return result
}
