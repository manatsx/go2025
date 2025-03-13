package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener valores de las variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	port := os.Getenv("PORT")

	// Imprimir valores
	fmt.Println("📌 Configuración cargada desde .env:")
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASS:", dbPass)
	fmt.Println("PORT:", port)

	// Leer archivo .env personalizado sin cargarlo en os.Getenv
	myEnv, err := godotenv.Read("otherFolder/.env")
	if err != nil {
		fmt.Println("Error leyendo el archivo .env:", err)
	} else {
		fmt.Println("Variables leídas (sin cargar en el entorno):", myEnv)
	}

	// Cargar archivo en el entorno para poder acceder con os.Getenv
	if err := godotenv.Load("otherFolder/.env"); err != nil {
		log.Fatal("Error cargando el archivo .env:", err)
	}

	// Acceder a las variables desde el entorno del sistema
	fmt.Println("MY_ENV1:", os.Getenv("MY_ENV1"))
	fmt.Println("MY_ENV2:", os.Getenv("MY_ENV2"))
}
