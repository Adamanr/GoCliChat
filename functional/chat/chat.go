package Chats

import (
	Structs "CliChat/functional/struct"
	"fmt"
)

func Run(room *Structs.Room) {

	fmt.Printf("Вы вошли в комнату =%s=\n", room.Name)
	var message string
	for {
		fmt.Scanln(&message)
		fmt.Printf("%s# $s\n", room.Users, message)
		if message == "/leave" {
			return
		}
	}
	return
}
