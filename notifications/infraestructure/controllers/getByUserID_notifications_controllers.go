package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/JosephAntony37900/API-Hexagonal-1-Consumidor/notifications/application"
)

type GetNotificationsByUserController struct {
    useCase *application.GetNotificationsByUser
}

func NewGetNotificationsByUserController(useCase *application.GetNotificationsByUser) *GetNotificationsByUserController {
    return &GetNotificationsByUserController{useCase: useCase}
}

// GetNotificationsByUser maneja la solicitud para obtener notificaciones de un usuario
func (c *GetNotificationsByUserController) GetNotificationsByUser(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("usuario_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
        return
    }

    notifications, err := c.useCase.Run(userID)
    if err != nil {
        http.Error(w, "Error al obtener notificaciones", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(notifications)
}
