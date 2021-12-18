package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("on server: %s", http.ListenAndServe(":80", nil))
}
