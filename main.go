package main

import (
	"flag"
	"os"

	"github.com/nalehrma/chat/lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the provided ip")
	//reference to isHost is because flag.parse assigns a new value, we dont want a copy
	flag.Parse()
	//flag does same thing as os.args but with less work

	if isHost {
		// go run main.go -listen <ip>
		//os.Args stores the program name in the first indice, -listen in [2] providing the IP
		connIP := os.Args[2]
		//need to pass var to a function
		lib.RunHost(connIP)
		//fmt.Println("is host")
	} else {
		//go run main.go <ip>
		connIP := os.Args[1]
		//need to pass var to a function
		lib.RunGuest(connIP)
		//fmt.Println("is guest")
	}
}
