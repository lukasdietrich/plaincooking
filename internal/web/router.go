package web

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

type handlerFunc func(echo.Context, models.Querier) error

func NewRouter(db *sql.DB, querier models.Querier) http.Handler {
	r := echo.New()
	r.Use(middleware.Recover())

	tx := makeTransactionWrapper(db)

	api := r.Group("/api")
	api.GET("/recipes", tx(listRecipes))
	api.POST("/recipes", tx(createRecipe))
	api.GET("/recipes/:recipeId", tx(readRecipe))
	api.PUT("/recipes/:recipeId", tx(updateRecipe))
	api.DELETE("/recipes/:recipeId", tx(deleteRecipe))

	return r
}

func makeTransactionWrapper(db *sql.DB) func(handlerFunc) echo.HandlerFunc {
	return func(handler handlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			tx, err := db.BeginTx(ctx.Request().Context(), nil)
			if err != nil {
				return fmt.Errorf("could not begin transaction: %w", err)
			}

			defer tx.Rollback()

			if err := handler(ctx, models.New(tx)); err != nil {
				return err
			}

			return tx.Commit()
		}
	}
}
