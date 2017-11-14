package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/common/messaging"

	"github.com/joaoaneto/radiup/cycle/service"
	"github.com/joaoaneto/radiup/cycle/service/spotify"
)

func main() {

	fmt.Print("Cycle Server")

	//service.Start()

	initializeMessaging("amqp://guest:guest@localhost:5672")

	service.StartServer("6969")

}

func initializeMessaging(amqp_server_url string) {

	spotify.MessagingClient = &messaging.MessagingClient{}
	spotify.MessagingClient.ConnectToBroker(amqp_server_url)

}
