package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() http.Handler {
	r := echo.New()
	r.Use(middleware.Recover())

	return r
}
