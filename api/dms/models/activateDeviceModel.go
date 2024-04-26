package models

type ActivateDeviceModel struct {
	DeviceModel     string `json:"device_model"`
	DeviceBoardType string `json:"device_board_type"`
	SoftwareVersion string `json:"software_version"`
	CodeBaseUrl     string `json:"code_base_url"`
}
