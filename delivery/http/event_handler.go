package delivery

import (
	"net/http"
	"strconv"

	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"github.com/firstaadi-dev/narauction-clean-arch/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EventHandler struct {
	EventUsecase domain.EventUsecase
}

func (h EventHandler) FindEvent(c echo.Context) error {
	event, err := h.EventUsecase.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(200, event)
}

func (h EventHandler) GetEventById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	event, err := h.EventUsecase.GetById(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(200, event)
}

func (h EventHandler) UpcomingEvent(c echo.Context) error {
	event, err := h.EventUsecase.GetUpcoming()
	if err != nil {
		return err
	}
	return c.JSON(200, event)
}

func (h EventHandler) AddEvent(c echo.Context) error {
	var event domain.Event
	err := c.Bind(&event)
	if err != nil {
		return err
	}
	err = h.EventUsecase.Store(event)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")

}

func (h EventHandler) EditEvent(c echo.Context) error {
	var event domain.Event
	err := c.Bind(&event)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.EventUsecase.Update(uint(id), event)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)

}
func (h EventHandler) DeleteEvent(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.EventUsecase.Delete(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
func NewEventHandler(r *echo.Echo, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUsecase: us,
	}

	r.GET("/event", handler.FindEvent)
	r.GET("/event/:id", handler.GetEventById)
	r.GET("/event/upcoming", handler.UpcomingEvent)
	r.POST("/event", handler.AddEvent, middleware.BasicAuth(utils.BasicAuth))
	r.PUT("/event/:id", handler.EditEvent, middleware.BasicAuth(utils.BasicAuth))
	r.DELETE("/event/:id", handler.DeleteEvent, middleware.BasicAuth(utils.BasicAuth))

}
