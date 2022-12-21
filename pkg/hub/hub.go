package hub

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const maxMessageSize = 512

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Id int

	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[int]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

type emptyInterface interface{}

type WSMessage struct {
	Entity string      `json:"entity"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[int]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client.Id] = client
			fmt.Printf("clients map: &%v\n", h.clients)
		case client := <-h.Unregister:
			if _, ok := h.clients[client.Id]; ok {
				delete(h.clients, client.Id)
				close(client.Send)
				fmt.Printf("clients map after unregister: &%v\n", h.clients)
			}
			// case message := <-h.broadcast:
			// 	for id, client := range h.clients {
			// 		select {
			// 		case client.Send <- message:
			// 		default:
			// 			close(client.Send)
			// 			delete(h.clients, id)
			// 		}
			// 	}
		}
	}
}
func (h *Hub) SendMessage(channelId int, message WSMessage) bool {
	fmt.Print("send message")
	var conn = h.clients[channelId].Conn
	err := conn.WriteJSON(message)
	return err == nil
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
}
