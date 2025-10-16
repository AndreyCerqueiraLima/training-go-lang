package server

import (
	"bjj-system/internal/presentation/api"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.POST("/products", api.CreateProduct)
	e.GET("/products", api.GetProducts)
	e.Logger.Fatal(e.Start(":8080"))
}
