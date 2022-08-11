package Structs

type User struct {
	Id       int
	Login    string
	Password string
	Level    int
	Rooms    Room
}

type Message struct {
	Message     string
	User        string
	TimeCreated string
}

type Room struct {
	Id       int
	Name     string
	Users    []map[string]User
	Password string
	Message  []Message
}

type Chat struct {
	Rooms []Room
}
