package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

type Client struct {
	Connection *websocket.Conn
	Send       chan []byte
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

func WSRoute(rg *gin.RouterGroup, dbConnection *gorm.DB) {
	rg.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		client := &Client{
			Connection: conn,
			Send:       make(chan []byte),
		}

		conn.WriteMessage(websocket.TextMessage, []byte("Welcome to Prodev Media WS!"))
		conn.WriteMessage(websocket.TextMessage, []byte("You are connected to the server!"))
		// send many client connected
		conn.WriteMessage(websocket.TextMessage, []byte("There are "+fmt.Sprint(len(clients))+" clients connected"))

		clients[client] = true

		go handleWebSocketConnection(client)
	})

	go handleBroadcast()
}

func handleWebSocketConnection(client *Client) {
	defer func() {
		client.Connection.Close()
		delete(clients, client)
	}()

	for {
		_, message, err := client.Connection.ReadMessage()
		if err != nil {
			client.Connection.Close()
			delete(clients, client)
			break
		}

		broadcast <- message
	}
}

func handleBroadcast() {
	for {
		message := <-broadcast

		for client := range clients {
			select {
			case client.Send <- message:
			default:
				err := client.Connection.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					client.Connection.Close()
					delete(clients, client)
				}
			}
		}
	}
}
