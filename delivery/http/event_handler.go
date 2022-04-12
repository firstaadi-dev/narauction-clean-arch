package delivery

import (
	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"github.com/labstack/echo/v4"
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
func NewEventHandler(r *echo.Echo, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUsecase: us,
	}

	r.GET("/event", handler.FindEvent)
}
