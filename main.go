package main

import (
	"ws/config"
	"ws/websocket"
)

func main() {
	config.ConnectDB()

	hub := websocket.NewHub()

	wsHandler := websocket.NewHandler(hub)

	go hub.Run()


	websocket.InitRoute(wsHandler)
}