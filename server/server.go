package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

/*SocketServer as a object of a socket client*/
type SocketServer struct {
	Port string
	conn net.Conn
}

/*Init : initialize socket server and start it*/
func (server *SocketServer) Init() {
	ln, err := net.Listen("tcp", server.Port)
	errorCheck(err)

	defer ln.Close()

	log.Printf("Server is listening on port %v", server.Port)
	for {
		conn, err := ln.Accept()
		errorCheck(err)
		server.conn = conn
		defer conn.Close()
		defer server.conn.Close()
		log.Printf("Established a connection on %v [%v]", conn.LocalAddr().String(), conn)

		server.startConnection()
	}
}

func (server *SocketServer) startConnection() {
	for {
		message, err := bufio.NewReader(server.conn).ReadString('\n')
		if err != nil && err == io.EOF {
			server.closeConnection()
			break
		}
		errorCheck(err)

		message = strings.TrimSpace(message)

		if message == "close" {
			server.closeConnection()
			break
		}

		log.Printf("Message received: %v\n", message)

		newMessage := strings.ToUpper(message)

		log.Printf("Formatted message: %v\n", newMessage)

		server.conn.Write([]byte(newMessage + "\n"))
	}
}

func (server *SocketServer) closeConnection() {
	log.Printf("Connection ended %v [%v]", server.conn.LocalAddr().String(), server.conn)
	server.conn.Close()
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
