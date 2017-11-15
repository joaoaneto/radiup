package service

import (
	"log"
	"net/http"
)

func StartServer(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Start HTTP Server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port: " + port)
		log.Println("Error: " + err.Error())
	}

}