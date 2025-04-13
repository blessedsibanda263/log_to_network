package examples

import (
	"log"
	"net"
)

func NetworkLogClient() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("failed to connect to localhost:1902")
	}
	defer conn.Close()
	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("this is a regular message")
	logger.Panicln("this is a panic")
}
