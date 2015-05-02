package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var connections = map[string]net.Conn{}

func main() {
	var name string
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("error")
	}

	for {
		conn, err := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		name = strings.TrimSpace(string(message))
		connections[name] = conn //associate a name with the connection
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(name + " connected to server")
		go MessageListener(conn, name)
	}
}

func MessageListener(conn net.Conn, sender string) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		var mess string = string(message)
		mess = strings.TrimSpace(mess)
		if mess == "!close" {
			break
		}
		fmt.Println("mesage received: ", mess)
		go sendAll(mess, sender)
	}
	fmt.Println(sender + " disconnected from server")
	delete(connections, sender)
	go sendAll(sender+" disconnected from server", "server")
	conn.Close()
}

func sendAll(message string, name string) {
	for k, conn := range connections {
		message = strings.TrimSpace(message)
		k = strings.TrimSpace(k)
		name = strings.TrimSpace(name) //trim spac to get rd of any new line chaacters
		conn.Write([]byte(name + ": " + message + "\n"))
	}
}
