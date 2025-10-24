package handlers

import (
	"github.com/labstack/echo/v4"
)

func (gsi *GortenerServiceInterface) AddURL(e echo.Context) error {
	return e.String(200, "get url")
}
