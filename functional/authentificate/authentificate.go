package authentificate

import (
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
	answer string
)

func CliClear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Auth() bool {
	CliClear()
	for {
		fmt.Print("Вы уже есть в чате? (y/n)\t")
		fmt.Scanln(&answer)

		switch answer {
		case "y":
			if authenticate() {
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
	var login, password, second_password string
	fmt.Println("Введите логин: ")
	fmt.Scanf("%s", &login)
	fmt.Println("Введите пароль: ")
	fmt.Scanf("%s", &password)

	for i := 0; i < 3; i++ {
		fmt.Println("Введите повторно пароль: ")
		fmt.Scanf("%s", &second_password)
		if login == loginP && password == second_password {
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

func authenticate() bool {
	var login, password string

	for i := 0; i < 3; i++ {
		fmt.Print("Введите логин: ")
		fmt.Scanf("%s", &login)
		fmt.Print("Введите пароль: ")
		fmt.Scanf("%s", &password)
		if login == loginP && password == passP {
			fmt.Println("Вы авторизованы!")
			return true
		}
		fmt.Printf("Введены неверные данные! Попытка %d\n", i+1)
	}
	return false
}
