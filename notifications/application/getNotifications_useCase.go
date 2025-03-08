package application

import (
	"time"

	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/repository"
)

type GetNotificationsByUser struct {
	repo repository.NotificationRepository
}

func NewGetNotificationsByUser(repo repository.NotificationRepository) *GetNotificationsByUser {
	return &GetNotificationsByUser{repo: repo}
}

func (g *GetNotificationsByUser) Run(usuarioID int, lastTimestamp time.Time) ([]entities.Notification, error) {
	// Obtener las notificaciones del usuario
	notifications, err := g.repo.FindByUserID(usuarioID)
	if err != nil {
		return nil, err
	}

	// Filtrar notificaciones basándose en el last_timestamp
	var newNotifications []entities.Notification
	for _, notification := range notifications {
		if notification.CreatedAt.After(lastTimestamp) {
			newNotifications = append(newNotifications, notification)
		}
	}

	return newNotifications, nil
}