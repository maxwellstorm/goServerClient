package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter IP address of server")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	conn, err := net.Dial("tcp", userInput+":1234")
	// handle error
	if err != nil {
		fmt.Println("an eror has occured")
		os.Exit(1)
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
