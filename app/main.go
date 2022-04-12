package main

import (
	"github.com/firstaadi-dev/narauction-clean-arch/config"
	delivery "github.com/firstaadi-dev/narauction-clean-arch/delivery/http"
	repository "github.com/firstaadi-dev/narauction-clean-arch/repository/psql"
	"github.com/firstaadi-dev/narauction-clean-arch/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.DatabaseConnect()
	e := echo.New()

	barangRepo := repository.NewPsqlBarangRepository(db)
	barangUcase := usecase.NewBarangUsecase(barangRepo)

	eventRepo := repository.NewPsqlEventRepository(db)
	eventUcase := usecase.NewEventUsecase(eventRepo)

	delivery.NewBarangHandler(e, barangUcase)
	delivery.NewEventHandler(e, eventUcase)

	e.Start(":8080")
}
