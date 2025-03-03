package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"ARQ.HEX/src/core"
	"ARQ.HEX/src/internal/notas/infrastructure/dependencies/routes"
)

func main() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	r := gin.Default()

	deps := &routes.NotasDependencies{
		DB: db,
	}

	routes.RegisterNoteRoutes(r, deps) 

	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}