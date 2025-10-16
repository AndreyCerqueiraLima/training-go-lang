package server

import (
	"bjj-system/internal/controllers"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProductById)
	e.Logger.Fatal(e.Start(":8080"))
}
