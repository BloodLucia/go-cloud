package router

import (
	"net/http"

	v1 "github.com/3blank/cloud-competition-serve/api/v1"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func InitRouter(app *pocketbase.PocketBase) {

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		api := e.Router.Group("/api")
		api.AddRoute(echo.Route{
			Path:   "/game/:id",
			Method: http.MethodGet,
			Handler: func(c echo.Context) error {
				id := c.PathParam("id")
				r, err := app.Dao().FindRecordById("games", id)
				if err != nil {
					return c.JSON(http.StatusNotFound, map[string]any{
						"success": false,
						"message": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, r)
			},
		})
		// urg := api.Group("/user")

		// urg.Use(apis.RequireRecordAuth("users"))

		gag := api.Group("/game")
		grh := new(v1.GameRouteHandler)
		{
			gag.AddRoute(echo.Route{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: grh.Index(app),
			})
			gag.AddRoute(echo.Route{
				Method:  http.MethodGet,
				Path:    "/item/:id",
				Handler: grh.Get(app),
			})
			gag.AddRoute(echo.Route{
				Method: http.MethodGet,
				Path:   "/",
				Handler: func(c echo.Context) error {
					return c.String(http.StatusOK, "hello world")
				},
			})
		}

		return nil
	})
}
