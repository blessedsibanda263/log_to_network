package examples

import (
	"fmt"
	"log"
	"net"
	"time"
)

func UDPBasedLogging() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("failed to connect to localhost:1902")
	}
	defer conn.Close()
	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	// logger.Panicln("this is a panic.")
	count := 0
	for count < 100 {
		fmt.Println("count = ", count)
		count += 1
		logger.Println("this is a regular message.", count)
		time.Sleep(500 * time.Millisecond)
	}
}
