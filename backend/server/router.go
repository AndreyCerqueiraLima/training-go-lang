package server

import (
	"bjj-system/internal/presentation/api"
	"bjj-system/internal/repository"
	"bjj-system/internal/service"
	"bjj-system/pkg/db_driver"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	db := db_driver.GetInstance()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(&productRepository)
	productHandler := api.NewProductHandler(productService)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProductById)

	e.Logger.Fatal(e.Start(":8080"))
}
