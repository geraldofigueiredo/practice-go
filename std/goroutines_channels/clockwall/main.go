package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	var port string

	if args[0] != "--port" {
		log.Fatal("unable to find --port argument")
	}
	port = args[1]

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Linstening on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
