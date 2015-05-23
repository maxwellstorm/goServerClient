package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("please enter IP address and port of server")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		conn1, err := net.Dial("tcp", userInput)
		// handle error
		if err != nil {
			fmt.Println("error connecting to server")
		} else {
			conn = conn1
			break
		}
	}
	fmt.Println("enter a client name")
	name, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, name+"\n")
	name = strings.TrimSpace(name)
	go messListener(conn, name)

	for {
		userInput, _ := reader.ReadString('\n') //wait until user presses enter
		_, err := fmt.Fprintf(conn, userInput+"\n")
		if err != nil {
			fmt.Println("disconnected from server")
			break
		}
		if strings.TrimSpace(userInput) == "!close" {
			fmt.Println("goodbye")
			os.Exit(0)
		}
	}
}

func messListener(conn net.Conn, name string) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		if strings.Split(message, ":")[0] != name {
			//fmt.Println(strings.Split(message, ":")[0])

			fmt.Print(message)
		}
	}
	fmt.Println("disconnected from server")
}
