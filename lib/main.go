package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//defining the port that is listening
const port = "8080"

//RunHost receives ip as argument from guest ip
func RunHost(ip string) {
	//getting the ip and adding the port into a string
	ipAndPort := ip + ":" + port
	//storing net listen tcp results from ipAndPort
	//net.listen returns listener and an error so assign it to the two vars
	//listener, listenErr := net.Listen("tcp", ipAndPort)

	//go returns errors to a variable which is checked for said error
	// removing listener to illustrate returning error
	//_, listenErr := net.Listen("tcp", ipAndPort)
	listener, listenErr := net.Listen("tcp", ipAndPort)
	//error displays if other readErr different than nil
	if listenErr != nil {
		//can use log.Fatal which saves time
		//fmt.Println("Error", listenErr)
		//os.Exit(1)
		log.Fatal("Error: ", listenErr)
	}
	//prints that host is awaiting connection for usability, printing IP
	fmt.Println("Listening on... ", ipAndPort)

	//listener acccepts connections, returning connection obj & error
	conn, acceptErr := listener.Accept()
	//error displays if other readErr different than nil
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	//prints that it accepted the connection
	fmt.Println("New connection accepted")

	//creating loop so connection doesnt close after first message
	//passing to function, creating a infinite for loop
	for {
		handleHost(conn)
	}
}

//handleHost is a function for the above for loop and readability
func handleHost(conn net.Conn) {
	//read message from connection from buf io package
	reader := bufio.NewReader(conn)
	//ReadString returns 'read text', and an error
	//The new line signifies the end of the transmission
	message, readErr := reader.ReadString('\n')
	//error displays if other readErr different than nil
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	//host reads message from the connection and prints to console
	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}
	fmt.Fprint(conn, replyMessage)
}

//RunGuest takes destination ip and connects to it
func RunGuest(ip string) {
	//assigning ip and port to variable
	ipAndPort := ip + ":" + port
	//using dial for the guest to 'dial' in via TCP and ipAndPort info
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}

	for {
		handleGuest(conn)
	}

}

//handleGuest is a function for the above loop and readability
func handleGuest(conn net.Conn) {
	//displays send message prompt to guest
	fmt.Print("Send Message: ")

	//reading guest message identifying standard input
	reader := bufio.NewReader(os.Stdin)
	//new line signifies end of message
	message, readErr := reader.ReadString('\n')
	//error displays if other readErr different than nil
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}
	fmt.Println("Message received: ", replyMessage)
}
