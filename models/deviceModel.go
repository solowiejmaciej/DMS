package models

import "time"

type DeviceModel struct {
	DeviceId  string    `json:"device_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
