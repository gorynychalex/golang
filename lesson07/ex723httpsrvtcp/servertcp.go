package main

import (
	"fmt"
	// "log"
	"net"
)

func main() {

	message := "Hello my friend!"

	listener, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go func(con net.Conn) {
			conn.Write([]byte(message))
			defer conn.Close()	
			}(conn)
	
	}
}