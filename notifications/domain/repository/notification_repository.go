package repository

import "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"

type NotificationRepository interface {
	Save(notification entities.Notification) error
	FindByUserID(usuario_id int) ([]entities.Notification, error)
}