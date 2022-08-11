package main

import (
	"CliChat/functional/authentificate"
	Chats "CliChat/functional/chat"
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
		"/list - показать список комнат\n" +
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

func createRoom() Structs.Room {
	var password, name, answer string
	var room Structs.Room
	roomId := rand.Intn(100)

	fmt.Scan(&name)

	fmt.Print("Вы хотите задать пароль комнате? (y/n)\t")
	fmt.Scanf("%s", &answer)

	switch answer {
	case "y":
		fmt.Print("Введите пароль: ")
		fmt.Scanf("%s", &password)
		room = Structs.Room{Id: roomId, Name: name, Password: password}
		fmt.Printf("Комната [ Id: %d | Room Name: %s | password: %s ]\n", room.Id, room.Name, room.Password)
		break
	case "n":
		room = Structs.Room{Id: roomId, Name: name}
		fmt.Printf("Комната создана!\nId: %d | Room Name: %s\n", room.Id, room.Name)
		break
	}
	return room
}

func joinChat(chat Structs.Chat) {
	var roomId int
	var password string
	fmt.Scan(&roomId)
	for _, v := range chat.Rooms {
		if v.Id == roomId {
			if v.Password != "" {
				fmt.Println("В комнате установлен пароль")
				for i := 0; i < 3; i++ {
					fmt.Printf("Введите пароль (%d/3): ", i+1)
					fmt.Scanf("%s", &password)
					if i == 2 {
						fmt.Println("Вы не ввели пароль в три раза")
						return
					}
				}
			}
		}
	}
}

func main() {
	var answer string
	var chat Structs.Chat
	var room Structs.Room

	if authentificate.Auth() {
		go Chats.Server()
		authentificate.CliClear()
		fmt.Printf("\tДобро пожаловать в чат!\n")
		fmt.Printf("\tЗа помощью - /help\n\n")

		for {
			fmt.Println("###############Меню###############")
			fmt.Scanf("%s", &answer)
			switch answer {
			case "/help":
				printHelp()
				break
			case "/create":
				room = createRoom()
				chat.Rooms = append(chat.Rooms, room)
				break
			case "/join":
				joinChat(chat)
				break
			case "/list":
				printRooms(chat)
				break
			case "/exit":
				os.Exit(0)
				break
			default:
				fmt.Println("Неизвестная команда")
				break
			}
		}
	}

}
