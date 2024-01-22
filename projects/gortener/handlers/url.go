package handlers

import "github.com/labstack/echo/v4"

func GetURL(e echo.Context) error {
	return e.String(200, "get url")
}
