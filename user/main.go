package main

import (
	"fmt"

	"github.com/joaoaneto/radiup/user/repository"
	"github.com/joaoaneto/radiup/user/service"
)

var appName = "user-service"

func main() {

	mysql := repository.NewMySQLConfig()

	fmt.Println("Hello " + appName)

	service.Db = mysql
	service.StartServer("6767")

}
