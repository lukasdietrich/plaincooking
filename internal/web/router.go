package web

import (
	"log/slog"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/lukasdietrich/plaincooking/frontend"
	"github.com/lukasdietrich/plaincooking/internal/service"
)

func NewRouter(
	transactions *service.TransactionService,
	recipes *RecipeController,
	assets *AssetController,
) http.Handler {
	r := echo.New()
	r.Binder = new(binder)

	r.Use(middleware.Recover())
	r.Use(middleware.Gzip())
	r.Use(marshalApiError())
	r.Use(logError())
	r.Use(handleBusinessError())
	r.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "build",
		Filesystem: http.FS(frontend.Files()),
	}))

	api := r.Group("/api")
	api.Use(transactional(transactions))

	api.GET("/recipes", recipes.List)
	api.POST("/recipes", recipes.Create)
	api.GET("/recipes/:recipeId", recipes.Read)
	api.PUT("/recipes/:recipeId", recipes.Update)
	api.DELETE("/recipes/:recipeId", recipes.Delete)
	api.GET("/recipes/:recipeId/images", recipes.ListImages)
	api.POST("/recipes/:recipeId/images", recipes.UploadImage)

	api.GET("/assets/:assetId", assets.Download)

	logRoutes(r)
	return r
}

func logRoutes(r *echo.Echo) {
	routes := r.Routes()

	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Path == routes[j].Path {
			return routes[i].Method < routes[j].Method
		}

		return routes[i].Path <= routes[j].Path
	})

	for _, route := range routes {
		slog.Debug("registered route", slog.String("method", route.Method), slog.String("path", route.Path))
	}
}
