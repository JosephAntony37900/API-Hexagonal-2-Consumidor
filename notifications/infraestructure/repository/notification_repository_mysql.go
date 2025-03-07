package repository

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/domain/entities"
)

type NotificationRepositoryMySQL struct {
	db *sql.DB
}

func NewNotificationRepositoryMySQL(db *sql.DB) *NotificationRepositoryMySQL {
	return &NotificationRepositoryMySQL{db: db}
}

// Save almacena una notificación en la base de datos
func (r *NotificationRepositoryMySQL) Save(notification entities.Notification) error {
	query := "INSERT INTO notificaciones (Usuario_id, Mensaje) VALUES (?, ?)"
	_, err := r.db.Exec(query, notification.Usuario_id, notification.Mensaje)
	if err != nil {
		return fmt.Errorf("error al guardar la notificación: %v", err)
	}
	return nil
}

// FindByUserID obtiene las notificaciones de un usuario específico
func (r *NotificationRepositoryMySQL) FindByUserID(usuario_id int) ([]entities.Notification, error) {
	query := "SELECT Id, Usuario_id, Mensaje FROM notificaciones WHERE Usuario_id = ?"
	rows, err := r.db.Query(query, usuario_id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener notificaciones: %v", err)
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		if err := rows.Scan(&notification.Id, &notification.Usuario_id, &notification.Mensaje); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
