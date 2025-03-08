package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/controllers"
)

func NotificationRoutes(router *gin.Engine, createNotificationController *controllers.CreateNotificationController, getNotificationsController *controllers.GetNotificationsByUserController) {
    router.POST("/notifications", func(c *gin.Context) {
        createNotificationController.CreateNotification(c.Writer, c.Request)
    })

    router.GET("/notifications/:usuario_id", func(c *gin.Context) {
        getNotificationsController.GetNotificationsByUser(c)
    })
}