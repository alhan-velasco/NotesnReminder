package main

import (
	"log"

	"github.com/gin-gonic/gin"

	core "ARQ.HEX/src/core"

	notasDeps "ARQ.HEX/src/internal/notas/infrastructure/dependencies"
	recordatoriosDeps "ARQ.HEX/src/internal/recordatorio/infrastructure/dependencies"
)

func main() {
	// Conectar a la base de datos
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	r := gin.Default()

	notasDependencies := notasDeps.NewNotasDependencies(db)
	notasRouter := notasDependencies.GetRoutes()
	notasRouter.AttachRoutes(r)

	recordatoriosDependencies := recordatoriosDeps.NewRecordatoriosDependencies(db)
	recordatoriosRouter := recordatoriosDependencies.GetRoutes()
	recordatoriosRouter.AttachRoutes(r)

	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
