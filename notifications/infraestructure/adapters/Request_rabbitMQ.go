package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/repository"
)

var conn *amqp.Connection
var channel *amqp.Channel
type Order struct {
	Id         int    `json:"id"`
	Usuario_id int    `json:"usuario_id"`
	Producto   string `json:"producto"`
	Pais string `json:"pais"`
	Entidad_federativa string `json:"entidad_federativa"`
	Cp string `json:"cp"`
}

func InitRabbitMQ() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	username := os.Getenv("Name")
	password := os.Getenv("PasswordQueue")

	rabbitURL := fmt.Sprintf("amqp://%s:%s@98.85.106.157:5672/", username, password)
	conn, err = amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %s", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal en RabbitMQ: %s", err)
	}

	log.Println("Conectado a RabbitMQ exitosamente")
}

func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func ConsumeCreatedOrders(repo repository.NotificationRepository) {
	queue, err := channel.QueueDeclare(
		"created.order",
		true,            
		false,           
		false,           
		false,           
		nil,             
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola 'created.order': %s", err)
	}

	msgs, err := channel.Consume(
		queue.Name, 
		"",         
		true,       
		false,      
		false,      
		false,      
		nil,        
	)
	if err != nil {
		log.Fatalf("Error al registrar el consumidor: %s", err)
	}

	go func() {
		for d := range msgs {
			var order Order
			if err := json.Unmarshal(d.Body, &order); err != nil {
				log.Printf("Error al decodificar el mensaje: %s", err)
				continue
			}

			// Decidir aleatoriamente si el pedido es aceptado o rechazado
			rand.Seed(time.Now().UnixNano())
			decision := rand.Intn(2) // 0: Rechazado, 1: Aceptado

			var mensaje string
			var colaDestino string

			if decision == 1 {
				mensaje = fmt.Sprintf("Pedido aceptado para el producto: %s", order.Producto)
				colaDestino = "order.confirmed"
			} else {
				mensaje = fmt.Sprintf("Pedido rechazado para el producto: %s", order.Producto)
				colaDestino = "order.rejected"
			}

			notification := entities.Notification{
				Usuario_id: order.Usuario_id,
				Mensaje:    mensaje,
			}
			if err := repo.Save(notification); err != nil {
				log.Printf("Error al guardar la notificación: %s", err)
				continue
			}

			if err := PublishOrderDecision(order, colaDestino); err != nil {
				log.Printf("Error al publicar el mensaje en la cola %s: %s", colaDestino, err)
			}

			log.Printf("Notificación creada para el usuario %d: %s", order.Usuario_id, mensaje)
		}
	}()
}

func PublishOrderDecision(order Order, colaDestino string) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error al convertir la orden a JSON: %w", err)
	}

	err = channel.Publish(
		"",           
		colaDestino,  
		false,        
		false,       
		amqp.Publishing{
			ContentType: "application/json",
			Body:        orderJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}

	return nil
}