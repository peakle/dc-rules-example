package main

import (
	"fmt"
	"strings"
	"time"
)

var msg = fmt.Sprintf("123") // rule match example

func main() {
	s := strings.Replace(msg, "1", "", -1) //rule match example

	l := time.Now().Second() // rule match example
	fmt.Println(s, l)
}
