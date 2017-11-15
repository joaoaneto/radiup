package service

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func StartServer(port string) {

	r := NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	handler := c.Handler(r)

	http.Handle("/", handler)

	log.Println("Start HTTP Server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port: " + port)
		log.Println("Error: " + err.Error())
	}

}
