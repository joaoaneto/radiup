package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/playlist/service"
)

func main() {

	fmt.Print("Playlist Server")

	//service.Start()

	//initializeMessaging("amqp://guest:guest@localhost:5672")

	service.StartServer("6868")

}

/*
func initializeMessaging(amqp_server_url string) {

	spotify.MessagingClient = &messaging.MessagingClient{}
	spotify.MessagingClient.ConnectToBroker(amqp_server_url)

}
*/
