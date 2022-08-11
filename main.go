package main

import (
	"CliChat/functional/authentificate"
	Structs "CliChat/functional/struct"
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
)

func printHelp() {
	fmt.Println("Список команд:\n" +
		"/help - вывести список команд\n" +
		"/create - создать комнату\n" +
		"/join - присоединиться к комнате\n" +
		"/leave - покинуть комнату\n" +
		"/list - показать список комнат\n" +
		"/users - показать список пользователей в комнате\n" +
		"################################################")
}
func printRooms(chat Structs.Chat) {
	fmt.Println("Список комнат: ")
	if chat.Rooms != nil {
		for _, room := range chat.Rooms {
			fmt.Println(room.Id, " ", room.Name)
		}
		return
	}
	fmt.Println("-#-Нет комнат-#-")
}

func createRoom(chat Structs.Chat) Structs.Room {
	roomId := rand.Intn(100)

	fmt.Print("Введите название комнаты: ")
	name, _ := in.ReadString('\n')
	fmt.Print("Вы хотите задать пароль комнате? (y/n)\t")
	answer, _ := in.ReadString('\n')
	switch answer {
	case "y":
		fmt.Print("Введите пароль: ")
		password, _ := in.ReadString('\n')
		chat.Rooms = append(chat.Rooms, Structs.Room{Id: roomId, Name: name, Password: password})
		fmt.Printf("Комната создана!\nId: %d | Room Name: %s | password: %s\n", roomId, name, chat.Rooms[len(chat.Rooms)-1].Password)
		break
	case "n":
		chat.Rooms = append(chat.Rooms, Structs.Room{Id: roomId, Name: name})
		fmt.Printf("Комната создана!\nId: %d | Room Name: %s\n", roomId, name)
		break
	}
	fmt.Println("################################################")
	return chat.Rooms[len(chat.Rooms)-1]
}

func joinChat(chat Structs.Chat) {
	var roomId int
	var password string

	if chat.Rooms != nil {
		fmt.Println("Введите название комнаты: ")
		fmt.Scanln(&roomId)

		for _, room := range chat.Rooms {
			if room.Id == roomId {
				if room.Password != "" {
					fmt.Print("В комнате есть пароль! (3 попытки)\n")
					for i := 0; i < 3; i++ {
						fmt.Print("Введите пароль: ")
						password, _ = in.ReadString('\n')
						if password == room.Password {
							break
						}
						fmt.Println("Неверный пароль! Попробуйте еще раз!")
					}
				}
				//authentificate.CliClear()
				fmt.Println("Вы присоединились к комнате!")
				fmt.Println("В комнате ", room.Name, " пользователей: ", len(room.Users))
			}
			fmt.Println("Неправильный ID комнаты!")
		}
	}
	fmt.Println("Активных комнат нет")
}

func main() {
	authentificate.Auth()
	authentificate.CliClear()
	var chat Structs.Chat
	fmt.Printf("\tДобро пожаловать в чат!\n")
	fmt.Printf("\tЗа помощью - /help\n\n")
	for {
		answer, _ := in.ReadString('\n')
		switch answer {
		case "/help":
			printHelp()

		case "/create":
			chat.Rooms = append(chat.Rooms, createRoom(chat))
			break
		case "/join":
			joinChat(chat)
			break
		case "/leave":
			break
		case "/list":
			printRooms(chat)
			break
		case "/users":
			for _, room := range chat.Rooms {
				fmt.Println(room.Users)
			}
			break
		}
	}
}
