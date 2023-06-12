package v1

import (
	"fmt"
	"net/http"

	"github.com/3blank/cloud-competition-serve/model"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

type GameRequest struct {
}

type GameResponse struct {
	*models.Record
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
	Users     []string `json:"users"`
}

type GameRouteHandler struct{}

func (grh *GameRouteHandler) Index(pb *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var total int
		if err := pb.DB().Select("*").From("games").Row(&total); err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print(total)
		return c.String(http.StatusOK, "hello world")
	}
}

func (grh *GameRouteHandler) Get(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		gameModel := new(model.GameModel)
		r, err := app.Dao().FindRecordById(gameModel.TableName(), id)
		if err != nil {
			return apis.NewNotFoundError("没有找到资源，请检查请求是否正确", err.Error())
		}

		return c.JSON(http.StatusOK, r)
	}
}

func (grh *GameRouteHandler) AddUser(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		qp := c.QueryParams()
		gameId := qp.Get("gameId")
		// userId := qp.Get("userId")

		r, err := app.Dao().FindRecordById("games", gameId)

		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		userIds := r.Get("users").([]string)

		if len(userIds) > 0 {
			// userIds = append(userIds, userId)
			// r.Set("users", userIds)
			fmt.Print(userIds)
			return c.JSON(http.StatusOK, r)
		}

		return nil
	}
}
