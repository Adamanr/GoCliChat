package authentificate

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	loginP = "admin"
	passP  = "admin"
	in     = bufio.NewReader(os.Stdin)
)

func CliClear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Auth() bool {
	for {
		fmt.Print("Вы уже есть в чате? (y/n)\t")
		var answer string
		fmt.Scanln(&answer)
		switch answer {
		case "y":
			if authentificate() {
				CliClear()
				return true
			}
		case "n":
			if registration() {
				break
			}
		default:
			fmt.Println("Введен неверный ответ!")
		}
	}
}

func registration() bool {

	fmt.Println("Введите логин: ")
	login, _ := in.ReadString('\n')
	fmt.Println("Введите пароль: ")
	pass, _ := in.ReadString('\n')
	for i := 0; i < 3; i++ {
		fmt.Println("Введите повторно пароль: ")
		pass2, _ := in.ReadString('\n')
		if login == loginP && pass == pass2 {
			fmt.Println("Вы зарегистрированы!")
			return true
		}
		fmt.Println("Введены неверные данные!")
	}
	return false
}

func connectDb() mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("<MONGODB_URI>"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(databases)
	return *client
}

func authentificate() bool {
	for i := 0; i < 3; i++ {
		fmt.Println("Введите логин: ")
		login, _ := in.ReadString('\n')
		fmt.Println("Введите пароль: ")
		pass, _ := in.ReadString('\n')
		if login == loginP && pass == passP {
			fmt.Println("Вы авторизованы!")
			return true
		}
		fmt.Printf("Введены неверные данные! Попытка %d\n", i+1)
	}
	return false
}
