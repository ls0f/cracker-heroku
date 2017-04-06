package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lovedboy/cracker/cracker/logger"
	"github.com/lovedboy/cracker/cracker/proxy"
)


var g = logger.GetLogger()

func main() {

	var port, secret string
	port = os.Getenv("PORT")
	secret = os.Getenv("SECRET")
	if (port == ""){
		port = "80"
	}
	if secret == ""{
		secret = "123456"
	}
	addr := fmt.Sprintf(":%s", port)
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	logger.InitLogger(*debug)
	p := proxy.NewHttpProxy(addr, secret, false)
	p.Listen()

}
