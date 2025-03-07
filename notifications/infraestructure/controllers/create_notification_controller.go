package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/application"
)

type CreateNotificationController struct {
    useCase *application.CreateNotification
}

func NewCreateNotificationController(useCase *application.CreateNotification) *CreateNotificationController {
    return &CreateNotificationController{useCase: useCase}
}

// CreateNotification maneja la solicitud para crear una notificación
func (c *CreateNotificationController) CreateNotification(w http.ResponseWriter, r *http.Request) {
    var notification struct {
        Usuario_id int    `json:"usuario_id"`
        Mensaje    string `json:"mensaje"`
    }
    if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
        http.Error(w, "Error en los datos de entrada", http.StatusBadRequest)
        return
    }

    if err := c.useCase.Run(notification.Usuario_id, notification.Mensaje); err != nil {
        http.Error(w, "Error al guardar la notificación", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Notificación creada correctamente"})
}
