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

func (c *GetNotificationsByUserController) GetNotificationsByUser(ctx *gin.Context) {
	userIDStr := ctx.Param("usuario_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	lastTimestampStr := ctx.Query("last_timestamp")
	var lastTimestamp time.Time
	if lastTimestampStr != "" {
		lastTimestamp, err = time.Parse(time.RFC3339, lastTimestampStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de timestamp inválido"})
			return
		}
	}

	timeout := time.After(30 * time.Second) 
	tick := time.Tick(1 * time.Second)    

	for {
		select {
		case <-timeout:
			ctx.JSON(http.StatusOK, gin.H{"notifications": []entities.Notification{}, "last_timestamp": time.Now().Format(time.RFC3339)})
			return
		case <-tick:
			notifications, err := c.useCase.Run(userID, lastTimestamp)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener notificaciones"})
				return
			}

			if len(notifications) > 0 {
				ctx.JSON(http.StatusOK, gin.H{"notifications": notifications, "last_timestamp": time.Now().Format(time.RFC3339)})
				return
			}
		}
	}
}