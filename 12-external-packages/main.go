package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// UUID v1 (Basado en timestamp y MAC Address)
	idV1, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error generando UUID v1:", err)
	} else {
		fmt.Println("UUID v1:", idV1)
	}

	// UUID v3 (Basado en nombre + MD5)
	namespace := uuid.Must(uuid.NewRandom()) // Namespace aleatorio
	name := "example.com"
	idV3 := uuid.NewMD5(namespace, []byte(name))
	fmt.Println("UUID v3:", idV3)

	// UUID v4 (Aleatorio)
	idV4 := uuid.New()
	fmt.Println("UUID v4:", idV4)

	// UUID v5 (Basado en nombre + SHA-1)
	idV5 := uuid.NewSHA1(namespace, []byte(name))
	fmt.Println("UUID v5:", idV5)

	// Última versión (v1.6) - Sigue generando v4
	idLatest := uuid.New()
	fmt.Println("UUID v1.6:", idLatest)
	fmt.Println("Versión de UUID v1.6:", idLatest.Version())
}
