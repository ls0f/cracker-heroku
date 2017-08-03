package main

import (
	"fmt"
	"os"

	"github.com/ls0f/cracker"
)

func main() {

	var port, secret string
	port = os.Getenv("PORT")
	secret = os.Getenv("SECRET")
	if port == "" {
		port = "80"
	}
	if secret == "" {
		secret = "123456"
	}
	addr := fmt.Sprintf(":%s", port)
	p := cracker.NewHttpProxy(addr, secret, false)
	p.Listen()

}
