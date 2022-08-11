package Chats

import (
	"CliChat/functional/authentificate"
	Structs "CliChat/functional/struct"
	"fmt"
	"github.com/gin-gonic/gin"
	socket "github.com/gorilla/websocket"
	"sync"
)

var upgrader = socket.Upgrader{
	WriteBufferPool: &sync.Pool{},
	ReadBufferSize:  4048,
}

func Server() {
	r := gin.New()
	r.GET("/ws", func(c *gin.Context) {

	})
	r.GET("/rs", func(c *gin.Context) {

	})
	r.Run(":8080")
}

func Run(room Structs.Room) {

	authentificate.CliClear()
	fmt.Printf("Вы вошли в комнату =%s=\n", room.Name)
	var message string
	for _, v := range room.Message {
		fmt.Scanln(&message)
		fmt.Printf("%s# $s\n", v.User, v.Message)
		if message == "/leave" {
			return
		}
	}
	return
}
