package delivery

import (
	"net/http"
	"strconv"

	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type BarangHandler struct {
	BarangUsecase domain.BarangUsecase
}

func (h BarangHandler) FindBarangs(c echo.Context) error {
	barang, err := h.BarangUsecase.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(200, barang)
}

func (h BarangHandler) GetBarangById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	barang, err := h.BarangUsecase.GetById(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(200, barang)
}

func (h BarangHandler) GetBarangByEventId(c echo.Context) error {
	eventId, err := strconv.ParseUint(c.Param("eventId"), 10, 64)
	if err != nil {
		return err
	}
	barang, err := h.BarangUsecase.GetByEventId(uint(eventId))
	if err != nil {
		return err
	}
	return c.JSON(200, barang)
}

func (h BarangHandler) AddBarang(c echo.Context) error {
	var barang domain.Barang
	err := c.Bind(&barang)
	if err != nil {
		return err
	}
	err = h.BarangUsecase.Store(barang)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")
}

func (h BarangHandler) EditBarang(c echo.Context) error {
	var barang domain.Barang
	err := c.Bind(&barang)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.BarangUsecase.Update(uint(id), barang)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h BarangHandler) DeleteBarang(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.BarangUsecase.Delete(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func NewBarangHandler(r *echo.Echo, us domain.BarangUsecase) {
	handler := &BarangHandler{
		BarangUsecase: us,
	}
	r.GET("/barang", handler.FindBarangs)
	r.GET("/barang/:id", handler.GetBarangById)
	r.POST("/barang", handler.AddBarang)
	r.PUT("/barang/:id", handler.EditBarang)
	r.DELETE("/barang/:id", handler.DeleteBarang)
	r.GET("/event/barang/:eventId", handler.GetBarangByEventId)
}
