package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Escucha una conexión tcp en el puerto 8080
	li, err := net.Listen("tcp", ":8080")
	// Si existe algun error al escuchar el puerto, genera un log con el error y se cierra la conexión
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	// La conexión itera infinitamente, es decir, la conexión va a estar siempre abierta
	for {
		// Acepta conexiones entrantes
		conn, err := li.Accept()
		// Si existe un error, se imprime el mensaje y sigue la iteración
		if err != nil {
			log.Println(err)
			continue
		}

		// Se escribe un mensaje en la conexión
		// Envia respuestas al cliente:
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
