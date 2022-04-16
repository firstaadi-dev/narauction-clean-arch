package utils

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	godotenv.Load()
	PASS := os.Getenv("PASS")
	if username == "admin" && password == PASS {
		return true, nil
	}
	return false, nil
}
