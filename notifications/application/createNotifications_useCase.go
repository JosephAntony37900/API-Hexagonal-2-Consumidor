package application

import (
    "fmt"

    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/repository"
)

type CreateNotification struct {
    repo repository.NotificationRepository
}

func NewCreateNotification(repo repository.NotificationRepository) *CreateNotification {
    return &CreateNotification{repo: repo}
}

func (cn *CreateNotification) Run(usuario_id int, mensaje string) error {
    notification := entities.Notification{
        Usuario_id: usuario_id,
        Mensaje:    mensaje,
    }

    if err := cn.repo.Save(notification); err != nil {
        return fmt.Errorf("error al guardar la notificación: %w", err)
    }

    return nil
}
