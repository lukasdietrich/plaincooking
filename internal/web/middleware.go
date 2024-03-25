package web

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/lukasdietrich/plaincooking/internal/service"
)

func transactional(transactions *service.TransactionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			rctx := ctx.Request().Context()

			return transactions.Transactional(rctx, func(tctx context.Context) error {
				req := ctx.Request()
				req = req.WithContext(tctx)

				ctx.SetRequest(req)

				return next(ctx)
			})
		}
	}
}
