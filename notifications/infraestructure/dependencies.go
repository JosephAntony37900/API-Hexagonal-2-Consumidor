package infraestructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	notiApp "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/application"
	notiController "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/controllers"
    notiRepo "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/repository"
    notiRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/routes"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/adapters"
)

func InitNotificationsDependencies(Engine *gin.Engine, db *sql.DB){
	adapters.InitRabbitMQ()
	defer adapters.CloseRabbitMQ()

	notiRepository := notiRepo.NewNotificationRepositoryMySQL(db)

	go adapters.ConsumeCreatedOrders(notiRepository)
	createNoti := notiApp.NewCreateNotification(notiRepository)
	getByUserNoti := notiApp.NewGetNotificationsByUser(notiRepository)

	createNotiController := notiController.NewCreateNotificationController(createNoti)
	getByUserNotiController := notiController.NewGetNotificationsByUserController(getByUserNoti)

	notiRoutes.NotificationRoutes(Engine, createNotiController, getByUserNotiController)

}