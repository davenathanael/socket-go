package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/*SocketClient as a object of a socket client*/
type SocketClient struct {
	Host string
	conn net.Conn
}

/*Init : initialize socket connection and start it*/
func (client *SocketClient) Init() {
	conn, err := net.Dial("tcp", client.Host)
	client.conn = conn
	errorCheck(err)

	defer conn.Close()
	defer client.conn.Close()

	log.Printf("Established a connection on %v [%v]", conn.LocalAddr().String(), conn)

	client.startConnection()
}

func (client *SocketClient) startConnection() {
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send:")
		message, _ := reader.ReadString('\n')

		fmt.Fprintf(client.conn, message+"\n")

		if strings.TrimSpace(message) == "close" {
			client.conn.Close()
			break
		}

		reply, _ := bufio.NewReader(client.conn).ReadString('\n')
		fmt.Printf("Reply from server: %v\n", reply)
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
