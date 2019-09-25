package server

import (
	"log"
	"net/http"

	"service"
)

// StartServer ...
func StartServer(port string) {
	log.Println("server listening at" + port)

	r := service.NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
