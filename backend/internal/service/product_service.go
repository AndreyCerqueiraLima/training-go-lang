package service

import (
	"bjj-system/internal/model"
	"bjj-system/internal/model/dto"
	"bjj-system/internal/repository"
)

type ProductService struct {
	rep *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{rep: repo}
}

func (s *ProductService) CreateProduct(productIn dto.ProductIn) (uint32, error) {
	repository := *s.rep

	product := model.Product{Name: productIn.Name, Price: productIn.Price}

	if err := repository.Create(&product); err != nil {
		return 0, err
	}
	return product.Id, nil
}

func (s *ProductService) GetProducts() []dto.ProductOut {
	repository := *s.rep
	result := []dto.ProductOut{}
	products, err := repository.FindAll()
	if err != nil {
		return result
	}

	for _, product := range products {
		result = append(result, dto.ProductOut{Name: product.Name, Price: product.Price})
	}

	return result
}

func (s *ProductService) FindById(id int) (dto.ProductOut, error) {
	repository := *s.rep
	product, err := repository.FindById(id)

	result := dto.ProductOut{Name: product.Name, Price: product.Price}

	return result, err
}
