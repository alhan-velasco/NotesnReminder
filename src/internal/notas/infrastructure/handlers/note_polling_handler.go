package handlers

import (
	"ARQ.HEX/src/internal/notas/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var lastNoteID int = 0
var lastNoteCount int = 0
var lastCheckTime time.Time = time.Now()

func GetNewNote(service services.NoteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		notes, err := service.GetAllNotes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(notes) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No hay notas disponibles"})
			return
		}

		latestNote := notes[len(notes)-1]
		latestNoteID := latestNote.ID

		if latestNoteID > lastNoteID {
			lastNoteID = latestNoteID
			c.JSON(http.StatusOK, gin.H{
				"message": "Nueva nota encontrada",
				"note":    latestNote,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "No hay nuevas notas"})
		}
	}
}

func GetDeletedNotes(service services.NoteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		notes, err := service.GetAllNotes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(notes) < lastNoteCount {
			lastNoteCount = len(notes)
			c.JSON(http.StatusOK, gin.H{"message": "Una o más notas han sido eliminadas"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "No hay notas eliminadas"})
		}
	}
}

func WaitForNewNote(service services.NoteService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Timeout total (ej. 30 segundos)
        timeout := time.After(30 * time.Second)
        // Intervalo de verificación (ej. cada 2 segundos)
        ticker := time.NewTicker(2 * time.Second)
        defer ticker.Stop() // Aseguramos liberar recursos

        // Obtenemos el último ID de nota al inicio
        initialNotes, err := service.GetAllNotes()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        lastID := 0
        if len(initialNotes) > 0 {
            lastID = initialNotes[len(initialNotes)-1].ID
        }

        for {
            select {
            case <-timeout:
                // Timeout: No hubo cambios
                c.JSON(http.StatusRequestTimeout, gin.H{
                    "message": "No se encontraron notas nuevas después de 30 segundos",
                })
                return
            case <-ticker.C:
                // Verificamos notas actuales
                currentNotes, err := service.GetAllNotes()
                if err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                    return
                }

                if len(currentNotes) > 0 {
                    currentLastID := currentNotes[len(currentNotes)-1].ID
                    if currentLastID > lastID {
                        // ¡Nueva nota encontrada!
                        c.JSON(http.StatusOK, gin.H{
                            "message": "Nueva nota creada",
                            "note":    currentNotes[len(currentNotes)-1],
                        })
                        return
                    }
                }
            }
        }
    }
}

func WaitForNoteDeletion(service services.NoteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := time.After(30 * time.Second)
		ticker := time.NewTicker(5 * time.Second)

		initialNotes, _ := service.GetAllNotes()
		initialCount := len(initialNotes)

		for {
			select {
			case <-timeout:
				c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron eliminaciones en 30 segundos"})
				return
			case <-ticker.C:
				currentNotes, _ := service.GetAllNotes()
				if len(currentNotes) < initialCount {
					c.JSON(http.StatusOK, gin.H{"message": "Una nota ha sido eliminada"})
					return
				}
			}
		}
	}
}
