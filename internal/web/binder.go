package web

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type binder struct {
	echo.DefaultBinder
}

func (b *binder) Bind(i any, ctx echo.Context) error {
	if err := b.DefaultBinder.BindPathParams(ctx, i); err != nil {
		return err
	}

	var (
		req         = ctx.Request()
		contentType = req.Header.Get(echo.HeaderConnection)
	)

	switch req.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		if strings.HasPrefix(contentType, echo.MIMEApplicationJSON) {
			return b.DefaultBinder.BindBody(ctx, i)
		}

	case http.MethodGet, http.MethodDelete, http.MethodHead:
		return b.DefaultBinder.BindQueryParams(ctx, i)
	}

	return nil
}
