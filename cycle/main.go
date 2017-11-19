package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/cycle/service"
)

func main() {

	fmt.Print("Cycle Server")

	//service.Start()

	service.StartServer("6969")

}
