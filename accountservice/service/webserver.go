package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	log.Println("Start HTTP server on port ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("start HTTP server failed on port", port)
		log.Println("err ", err.Error())
	}
}
