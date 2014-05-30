package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/rover", roverHandler)
	http.Handle("/", http.FileServer(http.Dir("./webapp")))
	http.ListenAndServe(":7070", nil)
}

func roverHandler(rw http.ResponseWriter, req *http.Request) {
	cmd := req.FormValue("c")
	value, err := strconv.Atoi(req.FormValue("v"))
	if err != nil {
		log.Printf("Error converting v query parameter to int: %+v", err)
	}

	log.Printf("sending command: %v, value: %+v", cmd, value)
	err = arduinoDo(cmd, value)
	if err != nil {
		log.Printf("Error sending command: %s, value: %+v to arduino: %+v", cmd, value, err)
	}
}
