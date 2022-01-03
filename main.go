package main

import (
	"fmt"
	"strings"
)

var msg = fmt.Sprintf("123") // rule match example

func main() {
	s := strings.Replace(msg, "1", "", -1) //rule match example

	fmt.Println(s)
}
