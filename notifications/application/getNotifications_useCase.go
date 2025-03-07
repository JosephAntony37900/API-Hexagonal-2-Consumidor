package application

import (
    "fmt"

    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/repository"
)

type GetNotificationsByUser struct {
    repo repository.NotificationRepository
}

func NewGetNotificationsByUser(repo repository.NotificationRepository) *GetNotificationsByUser {
    return &GetNotificationsByUser{repo: repo}
}

func (gn *GetNotificationsByUser) Run(usuario_id int) ([]entities.Notification, error) {
    notifications, err := gn.repo.FindByUserID(usuario_id)
    if err != nil {
        return nil, fmt.Errorf("error al obtener notificaciones: %w", err)
    }

    return notifications, nil
}
