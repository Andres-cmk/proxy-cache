package main

import (
	"log"
	"net/http"
)

func main() {

	log.Print("Listening on port 8080...")
	_ = http.ListenAndServe(":8080", nil)

}
