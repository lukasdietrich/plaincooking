package web

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type binder struct {
	echo.DefaultBinder
}

func (b *binder) Bind(i any, ctx echo.Context) error {
	err := b.DefaultBinder.Bind(i, ctx)
	if errors.Is(err, echo.ErrUnsupportedMediaType) {
		return nil
	}

	return err
}
