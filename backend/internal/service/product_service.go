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

func (s *ProductService) CreateProduct(productIn dto.ProductIn) dto.ProductOut {
	repository := *s.p
	productOut := dto.ProductOut{}

	product := model.Product{Name: productIn.Name, Price: productIn.Price}

	if err := repository.Create(product); err != nil {
		return productOut
	}
	return productOut
}

func (s *ProductService) GetProducts() []dto.ProductOut {
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
