package entities

import "time"

type Notification struct {
	Id         int
	Usuario_id int
	Mensaje    string
	CreatedAt  time.Time
}