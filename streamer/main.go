package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/streamer/service"
)

func main() {

	fmt.Print("Streamer Server")

	service.StartServer("6868")

}
