package main

import (
	"fmt"
	"log"
	"net/http"
)

var _ = fmt.Sprintf("123") // rule match example

func main() {
	log.Printf("on server: %s", http.ListenAndServe(":80", nil))
}
