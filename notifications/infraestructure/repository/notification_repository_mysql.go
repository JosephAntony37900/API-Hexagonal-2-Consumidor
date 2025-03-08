package repository

import (
	"database/sql"
	"fmt"
	"time"

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
	query := "INSERT INTO notificaciones (Usuario_id, Mensaje, CreatedAt) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, notification.Usuario_id, notification.Mensaje, time.Now())
	if err != nil {
		return fmt.Errorf("error al guardar la notificación: %v", err)
	}
	return nil
}

// FindByUserID obtiene las notificaciones de un usuario específico
func (r *NotificationRepositoryMySQL) FindByUserID(usuarioID int) ([]entities.Notification, error) {
	query := "SELECT Id, Usuario_id, Mensaje, CreatedAt FROM notificaciones WHERE Usuario_id = ? ORDER BY CreatedAt DESC"
	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener notificaciones: %v", err)
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		var createdAt string
		if err := rows.Scan(&notification.Id, &notification.Usuario_id, &notification.Mensaje, &createdAt); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		// Convertir CreatedAt a time.Time
		notification.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt) // Ajusta el formato según tu base de datos
		if err != nil {
			return nil, fmt.Errorf("error al parsear CreatedAt: %v", err)
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}