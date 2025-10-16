package api

import (
	"bjj-system/internal/model/dto"
	"bjj-system/internal/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := h.validator.Struct(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id, err := h.ps.CreateProduct(*product)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ocorreu um erro ao processar a request :)")
	}

	return c.JSON(http.StatusOK, id)
}

func (h ProductHandler) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, h.ps.GetProducts())
}

func (h ProductHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")
	idConverted, _ := strconv.Atoi(id)
	product, _ := h.ps.FindById(idConverted)
	return c.JSON(http.StatusOK, product)
}
