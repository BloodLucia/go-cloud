package main

import (
	"log"

	"github.com/3blank/cloud-competition-serve/router"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	router.InitRouter(app)

	app.OnRecordAfterUpdateRequest("users").Add(func(e *core.RecordUpdateEvent) error {
		log.Println(e.Record.Get("score"))
		return nil
	})

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	api := e.Router.Group("/api")

	// 	userAPIGroup := api.Group("/user-api")
	// 	userAPIGroup.Use(apis.RequireRecordAuth("users"))

	// 	gameAPIGroup := userAPIGroup.Group("/game")
	// 	{
	// 		gameAPIGroup.AddRoute(echo.Route{
	// 			Method: http.MethodGet,
	// 			Path:   "/list",
	// 			Handler: func(c echo.Context) error {
	// 				return c.String(http.StatusOK, "/game/list")
	// 			},
	// 		})
	// 	}

	// 	// questionGroup := api.Group("/question")
	// 	// questionGroup.AddRoute(echo.Route{
	// 	// 	Method: http.MethodPost,
	// 	// 	Path:   "/submit",
	// 	// 	Handler: func(c echo.Context) error {
	// 	// 		q := &model.QuestionRequest{}
	// 	// 		err := c.Bind(q)
	// 	// 		if err != nil {
	// 	// 			return c.String(http.StatusBadRequest, err.Error())
	// 	// 		}
	// 	// 		r, err := app.Dao().FindRecordById("questions", q.ID)
	// 	// 		if err != nil {
	// 	// 			return c.JSON(http.StatusNotFound, map[string]any{
	// 	// 				"success": false,
	// 	// 				"message": "资源不存在",
	// 	// 			})
	// 	// 		}

	// 	// 		if isOk := r.GetString("answer") == q.Answer; !isOk {
	// 	// 			return c.JSON(http.StatusForbidden, map[string]any{
	// 	// 				"success": false,
	// 	// 				"message": "回答错误",
	// 	// 			})
	// 	// 		}

	// 	// 		return c.JSON(http.StatusOK, map[string]any{
	// 	// 			"success": true,
	// 	// 			"message": "回答正确",
	// 	// 		})
	// 	// 	},
	// 	// })

	// 	return nil
	// })

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
