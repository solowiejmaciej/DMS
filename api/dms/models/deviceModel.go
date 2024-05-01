package models

type DeviceModel struct {
	DeviceId        string `json:"device_id"`
	UserId          uint   `json:"user_id"`
	DeviceModel     string `json:"device_model"`
	DeviceBoardType string `json:"device_board_type"`
	SoftwareVersion string `json:"software_version"`
	CreatedAt       string `json:"created_at"`
	CodeBaseUrl     string `json:"code_base_url"`
}
