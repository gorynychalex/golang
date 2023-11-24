package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 11111})
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	message := "Привет от клиента UDP!"

	_, err = conn.Write([]byte(message))
	
	if err != nil {
		log.Println(err)
	}
}