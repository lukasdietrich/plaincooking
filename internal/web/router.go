package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

func NewRouter(querier models.Querier) http.Handler {
	r := echo.New()
	r.Use(middleware.Recover())

	return r
}
