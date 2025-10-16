package controllers

import (
	"bjj-system/internal/model/dto"
	"bjj-system/internal/repository"
	"bjj-system/internal/service"
	"bjj-system/pkg/db_driver"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()
var db = db_driver.GetInstance()
var productRepository = repository.NewProductRepository(db)
var productService = service.NewProductService(&productRepository)

func CreateProduct(c echo.Context) error {
	product := new(dto.ProductIn)

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := validate.Struct(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	productOut := productService.CreateProduct(*product)
	return c.JSON(http.StatusOK, productOut)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, productService.GetProducts())
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	product, _ := productService.FindById(idConverted)
	return c.JSON(http.StatusOK, product)
}
