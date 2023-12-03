package websocket

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(handler *Handler) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) (bool, error) {
			return origin == "http://localhost:3000", nil
		},
		MaxAge: int(12 * time.Hour),
	}))

	e.POST("/ws/createRoom", handler.CreateRoom)
	e.GET("/ws/joinRoom/:roomId", handler.JoinRoom)
	e.GET("/ws/getRooms", handler.GetRooms)
	e.GET("/ws/getClients/:roomId", handler.GetClients)
	e.Logger.Fatal(e.Start(":8000"))
}