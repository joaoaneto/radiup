package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/joaoaneto/radiup/common/messaging"

	"github.com/joaoaneto/radiup/user/model"
	"github.com/joaoaneto/radiup/user/repository"
	"github.com/joaoaneto/radiup/user/service"
)

var appName = "user-service"

var MessagingClient messaging.IMessagingClient

func main() {

	mysql := repository.NewMySQLConfig()

	fmt.Println("Hello " + appName)

	initializeMessaging("amqp://guest:guest@localhost:5672")

	service.Db = mysql
	service.StartServer("6767")

}

func onMessage(delivery amqp.Delivery) {
	fmt.Printf("Got a message: %v\n", string(delivery.Body))
	log.Print("teste")

	teste, _ := json.Marshal(string(delivery.Body))
	oAuthToken := model.SpotifyToken{}

	json.Unmarshal(delivery.Body, &oAuthToken)

	fmt.Println(teste)

}

func initializeMessaging(amqp_server_url string) error {

	MessagingClient = &messaging.MessagingClient{}
	MessagingClient.ConnectToBroker(amqp_server_url)

	err := MessagingClient.SubscribeToQueue("tokenQueue", appName, onMessage)
	if err != nil {
		panic("Could not start subscribe to tokenQueue.")
	}

	err = MessagingClient.Subscribe("topic_test", "topic", appName, onMessage)
	if err != nil {
		panic("Could not start subscribe to tokenQueue.")
	}

	return nil
}
