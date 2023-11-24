package main

import (
	// "fmt"
	"log"
	"net"
)

func main() {
	// Устанавливаем прослушивание порта
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	defer listen.Close()
	// Открываем порт
	conn, err := listen.Accept()
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Write([]byte("Привет мой друг!"))
	if err != nil {
    	log.Println(err)
	}
}