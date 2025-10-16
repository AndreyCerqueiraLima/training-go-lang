package controllers

import (
	"bjj-system/internal/model/dto"
	"bjj-system/internal/repository"
	"bjj-system/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()
var productRepository = repository.NewProductRepository()
var productService = service.NewProductService(&productRepository)

func CreateProduct(c echo.Context) error {
	product := dto.ProductIn{}

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := validate.Struct(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	productOut := productService.CreateProduct(product)
	return c.JSON(http.StatusOK, productOut)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, productService.GetProducts())
}
