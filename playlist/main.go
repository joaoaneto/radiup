package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/playlist/service"
	streamerService "github.com/joaoaneto/radiup/streamer/service"
)

func main() {

	fmt.Print("Playlist Server")

	streamerService.Start()

	//initializeMessaging("amqp://guest:guest@localhost:5672")

	service.StartServer("6969")

}

/*
func initializeMessaging(amqp_server_url string) {

	spotify.MessagingClient = &messaging.MessagingClient{}
	spotify.MessagingClient.ConnectToBroker(amqp_server_url)

}
*/
