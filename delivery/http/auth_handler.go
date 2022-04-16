package delivery

import (
	"net/http"

	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func (h AuthHandler) Login(c echo.Context) error {
	var userCreds domain.UserCredential
	err := c.Bind(&userCreds)
	if err != nil {
		return err
	}
	token, err := h.AuthUsecase.Login(userCreds)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, token)
}
func NewAuthHandler(r *echo.Echo, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: us,
	}
	r.POST("/login", handler.Login)
}
