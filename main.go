package main

import (
	"log"
	helpers "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/helpers"

	notiApp "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/application"
	notiController "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/controllers"
    notiRepo "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/repository"
    notiRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/routes"
    "github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/adapters"

)

func main () {
	db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error conectando a la BD: %v", err)
    }
    defer db.Close()

	// Inicializar RabbitMQ
	adapters.InitRabbitMQ()
	defer adapters.CloseRabbitMQ()

	notiRepository := notiRepo.NewNotificationRepositoryMySQL(db)

// Iniciar el consumidor de RabbitMQ
	go adapters.ConsumeCreatedOrders(notiRepository)
	createNoti := notiApp.NewCreateNotification(notiRepository)
	getByUserNoti := notiApp.NewGetNotificationsByUser(notiRepository)

	createNotiController := notiController.NewCreateNotificationController(createNoti)
	getByUserNotiController := notiController.NewGetNotificationsByUserController(getByUserNoti)

	// Configuración del enrutador de Gin
    r := gin.Default()

    // Configurar CORS
    r.Use(helpers.SetupCORS())

	//ruytas
	notiRoutes.NotificationRoutes(r, createNotiController, getByUserNotiController)

	 // Iniciar servidor
	 log.Println("Server started at :8081")
	 if err := r.Run(":8081"); err != nil {
		 log.Fatalf("Error starting server: %v", err)
	 }
}