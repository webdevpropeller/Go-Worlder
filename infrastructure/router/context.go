package router

import (
	"go_worlder_system/interface/controller"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type callFunc func(c controller.Context) error

func c(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(controller.Context))
	}
}

// Context ...
type Context struct {
	echo.Context
	upgrader websocket.Upgrader
}

// WebSocket ...
func (c Context) WebSocket() (*websocket.Conn, error) {
	return c.upgrader.Upgrade(c.Response(), c.Request(), nil)
}

// UserID ...
func (c Context) UserID() string {
	return c.Get(userIDKey).(string)
}
