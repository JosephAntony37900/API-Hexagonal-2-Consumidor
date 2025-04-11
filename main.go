package main

import (
	"log"
	helpers "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/helpers"
	init_noti "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/infraestructure/adapters"
    "github.com/gin-gonic/gin"

)

func main () {
	db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error conectando a la BD: %v", err)
    }
    defer db.Close()

    r := gin.Default()

    r.Use(helpers.SetupCORS())

    init_noti.InitNotificationsDependencies(r, db)

	defer adapters.CloseRabbitMQ()

	 log.Println("Server escuchandp :8081")
	 if err := r.Run(":8081"); err != nil {
		 log.Fatalf("Error iniciando el server: %v", err)
	 }
}