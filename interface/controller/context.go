package controller

import (
	"mime/multipart"

	"github.com/gorilla/websocket"
)

// Context ...
type Context interface {
	WebSocket() (*websocket.Conn, error)
	UserID() string
	String(code int, s string) error
	JSON(code int, i interface{}) error
	FormValue(name string) string
	QueryParam(name string) string
	FormFile(name string) (*multipart.FileHeader, error)
	Param(name string) string
	Bind(i interface{}) error
	IsWebSocket() bool
}
