package authorization

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

type UserAuthRouteHandler struct{}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uah *UserAuthRouteHandler) SignIn(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		sq := new(SignInRequest)
		if err := c.Bind(sq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{})
		}

		return c.JSON(http.StatusOK, sq)
	}
}

func (uah *UserAuthRouteHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		sq := new(SignInRequest)
		if err := c.Bind(sq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{})
		}

		return c.JSON(http.StatusOK, sq)
	}
}
