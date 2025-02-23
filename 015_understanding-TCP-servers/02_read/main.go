package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Escucha una conexión tcp en el puerto 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	// El server tcp siempre estara abierto
	for {
		// Acepta conexiones entrantes
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// procesar múltiples conexiones simultáneamente sin bloquear el flujo de ejecución del servidor
		go handle(conn)
	}
}

// func para procesar la conexión que se ha establecido con un cliente en el servidor TCP
// net.Conn representa una conexión de red abierta
func handle(conn net.Conn) {
	// lee los datos de la conexión
	scanner := bufio.NewScanner(conn)
	// se ejecuta mientras haya datos para leer en la conexión.
	for scanner.Scan() {
		// ln guarda el texto leido
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}

// An open connection blocks.
// The reader is reading from the open connection.
// How does the reader know when it should stop reading?
// Instructions: run this file, then in your browser go to:
// http://localhost:8080/
