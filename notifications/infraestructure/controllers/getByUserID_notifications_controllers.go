package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/application"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
	"github.com/gin-gonic/gin"
)

type GetNotificationsByUserController struct {
	useCase *application.GetNotificationsByUser
}

func NewGetNotificationsByUserController(useCase *application.GetNotificationsByUser) *GetNotificationsByUserController {
	return &GetNotificationsByUserController{useCase: useCase}
}

// GetNotificationsByUser maneja la solicitud para obtener notificaciones de un usuario
func (c *GetNotificationsByUserController) GetNotificationsByUser(ctx *gin.Context) {
	// Obtener el usuario_id de la URL
	userIDStr := ctx.Param("usuario_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Obtener el timestamp de la última notificación (si se proporciona)
	lastTimestampStr := ctx.Query("last_timestamp")
	var lastTimestamp time.Time
	if lastTimestampStr != "" {
		lastTimestamp, err = time.Parse(time.RFC3339, lastTimestampStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de timestamp inválido"})
			return
		}
	}

	// Long Polling: Esperar hasta que haya nuevas notificaciones
	timeout := time.After(30 * time.Second) // Tiempo máximo de espera
	tick := time.Tick(1 * time.Second)      // Intervalo de verificación

	for {
		select {
		case <-timeout:
			// Si no hay notificaciones después de 30 segundos, devolver un array vacío
			ctx.JSON(http.StatusOK, gin.H{"notifications": []entities.Notification{}, "last_timestamp": time.Now().Format(time.RFC3339)})
			return
		case <-tick:
			// Verificar si hay nuevas notificaciones
			notifications, err := c.useCase.Run(userID, lastTimestamp)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener notificaciones"})
				return
			}

			// Si hay notificaciones, devolverlas
			if len(notifications) > 0 {
				ctx.JSON(http.StatusOK, gin.H{"notifications": notifications, "last_timestamp": time.Now().Format(time.RFC3339)})
				return
			}
		}
	}
}