package transport

import (
	"blog/src/repository/chat"
	"blog/src/usecases"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type serviceChat struct {
	hub *chat.Hub
	svc usecases.ChatService
}

func NewChatService(e echo.Echo, chatService usecases.ChatService, middlewareFunc echo.MiddlewareFunc, hub *chat.Hub) {
	chatHTTP := &serviceChat{hub: hub}

	chat := e.Group("/chat")

	chat.GET("/ws", chatHTTP.serveWs)
	chat.GET("/", chatHTTP.serveHome)
}

func (s serviceChat) serveHome(c echo.Context) error {
	log.Println(c.Request().URL)

	http.ServeFile(c.Response(), c.Request(), "/home/yurchenkovr/go/src/blog/src/infrastructure/transport/home.html")

	return c.JSON(http.StatusOK, nil)
}

func (s serviceChat) serveWs(c echo.Context) error {
	conn, err := chat.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello Upgraded!"))
	if err != nil {
		log.Println("Error: 1")
	}
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error: 2")
	}
	fmt.Printf(":: %v", msg)
	client := chat.NewClient(s.hub, conn, make(chan []byte, 256), s.svc)
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump(c)

	return c.JSON(http.StatusOK, nil)
}
