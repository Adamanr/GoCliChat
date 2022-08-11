package Structs

type Message struct {
	Message     string
	User        string
	TimeCreated string
}

type Room struct {
	Id       int
	Name     string
	Users    []map[string]string
	Password string
	Message  []Message
}

type Chat struct {
	Rooms []Room
}
