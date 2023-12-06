package websocket

import (
	"net/http"
	"ws/config"
	"ws/model"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
}

func (h *Handler) CreateRoom(c echo.Context) error {
	var req CreateRoomReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error binding request",
		})
	}
	room := model.Rooms{
		ID: req.ID,
		Name: req.Name,
		PhotoProfile: req.PhotoProfile,
	}
	config.DB.Create(&room)
	h.hub.Rooms[room.ID] = &Room{
		ID:      room.ID,
		Name:    room.Name,
		Client: make(map[string]*Client),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "room created successfully",
	})
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error binding request",
		})
		
	}

	roomID := c.Param("roomId")
	clientID := c.QueryParam("userId")
	username := c.QueryParam("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}
	

	h.hub.Register <- cl

	go cl.writeMessage(roomID)
	cl.readMessage(h.hub)

	return nil
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c echo.Context) error {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c echo.Context) error {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Client {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return c.JSON(http.StatusOK, clients)
}

func (h *Handler) GetChats(c echo.Context) error {
	roomID := c.Param("roomId")
	var messages []model.Message
	config.DB.Where("room_id = ?", roomID).Find(&messages)

	return c.JSON(http.StatusOK, messages)
}