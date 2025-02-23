package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Dial se usa para crear una conexión de red, permitiendo que un cliente se conecte a un servidor en una dirección y red específicas.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
