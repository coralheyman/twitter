package main

import (
	"log"

	"github.com/coralheyman/twitter/bd"
	"github.com/coralheyman/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
