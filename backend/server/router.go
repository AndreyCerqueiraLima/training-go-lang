package server

import (
	"bjj-system/internal/presentation/api"
	"bjj-system/pkg/db_driver"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	db := db_driver.GetInstance()
	defer db.Close()

	api.ExecProductHandler(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
