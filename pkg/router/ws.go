package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tasktracker-api/pkg/hub"

	"github.com/gorilla/websocket"
)

type wsAuthMessage struct {
	Bearer  string `json:"bearer"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (router *Router) WSHandler(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print(err)
		return
	}
	var m wsAuthMessage
	var userId int
	var client *hub.Client
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		err = json.Unmarshal(p, &m)
		if err != nil {
			return
		}
		// fmt.Printf("message: %+v\n", m)
		claims, ok := router.GetClaimsFromToken(m.Bearer)
		if ok != nil {
			return
		}
		userIDString := claims["sub"].(string)
		// fmt.Printf("user: %+v\n", userIDString)
		userId, err = strconv.Atoi(userIDString)
		if err != nil {
			return
		}
		client = &hub.Client{Id: userId, Hub: h, Conn: conn, Send: make(chan []byte, 256)}
		client.Hub.Register <- client
	}
	go client.ReadPump()
}
