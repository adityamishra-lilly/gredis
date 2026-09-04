package main

import (
	"flag"
	"log"

	"github.com/adityamishra-lilly/gredis/config"
	"github.com/adityamishra-lilly/gredis/server"
)

func setupFlags() {
	flag.StringVar(
	&config.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.Port, "port", 7379, "port for the server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println()
	server.StartServer()
}