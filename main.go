package main

import (
	"os"
	"tcp-socket/client"
	"tcp-socket/server"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		startServer()
		return
	}

	if args[1] == "client" {
		if len(args) == 2 {
			startClient("localhost:8080")
		} else {
			startClient(args[2])
		}
	} else {
		startServer()
	}
}

func startClient(host string) {
	client := client.SocketClient{Host: host}
	client.Init()
}

func startServer() {
	server := server.SocketServer{Port: ":8080"}
	server.Init()
}
