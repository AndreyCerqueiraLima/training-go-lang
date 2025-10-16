package api

import (
	"bjj-system/internal/model/dto"
	"bjj-system/internal/repository"
	"bjj-system/internal/service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ExecProductHandler(e *echo.Echo, db *sql.DB) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(&productRepository)
	productHandler := NewProductHandler(productService)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProductById)
	e.DELETE("/products/:id", productHandler.DeleteProduct)
}

type ProductHandler struct {
	validator *validator.Validate
	ps        *service.ProductService
}

func NewProductHandler(ps *service.ProductService) ProductHandler {
	var validate = validator.New()
	return ProductHandler{ps: ps, validator: validate}
}

func (h ProductHandler) CreateProduct(c echo.Context) error {
	product := new(dto.ProductIn)

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{"error": "Invalid JSON"})
	}

	if err := h.validator.Struct(product); err != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{"error": err.Error()})
	}

	id, err := h.ps.CreateProduct(*product)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ocorreu um erro ao processar a request ")
	}

	return c.JSON(http.StatusOK, id)
}

func (h ProductHandler) GetProducts(c echo.Context) error {
	products := h.ps.GetProducts()
	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, products)
	}
	return c.JSON(http.StatusOK, h.ps.GetProducts())
}

func (h ProductHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	product, _ := h.ps.FindById(idConverted)
	return c.JSON(http.StatusOK, product)
}

func (h ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	deleted, err := h.ps.DeleteProduct(idConverted)

	if err != nil {
		return c.JSON(http.StatusNotFound, idConverted)
	}

	if !deleted {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	return c.JSON(http.StatusNoContent, true)
}
