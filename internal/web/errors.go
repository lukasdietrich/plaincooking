package web

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-sqlite3"
	"github.com/rs/xid"

	"github.com/lukasdietrich/plaincooking/internal/service"
)

func logError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)
			if err != nil {
				slog.Warn("request error",
					slog.String("url", ctx.Request().RequestURI),
					slog.String("err", err.Error()))
			}

			return err
		}
	}
}

func handleBusinessError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := next(ctx); err != nil {
				return mapBusinessError(err)
			}

			return nil
		}
	}
}

func mapBusinessError(err error) error {
	if errors.Is(err, xid.ErrInvalidID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id").SetInternal(err)
	}

	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "resource not found").SetInternal(err)
	}

	if errors.Is(err, service.ErrInvalidRecipe) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid recipe").SetInternal(err)
	}

	if sqliteErr, ok := errorAs[sqlite3.Error](err); ok {
		return mapSqliteError(sqliteErr)
	}

	return err
}

func mapSqliteError(err sqlite3.Error) error {
	if err.Code == sqlite3.ErrConstraint {
		return echo.NewHTTPError(http.StatusConflict, "constraint violation").SetInternal(err)
	}

	return err
}

func errorAs[E error](err error) (E, bool) {
	var errTyped E
	ok := errors.As(err, &errTyped)

	return errTyped, ok
}
