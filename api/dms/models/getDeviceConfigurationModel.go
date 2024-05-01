package models

type GetDeviceConfigurationModel struct {
	MqttBrokerPort int    `json:"mqtt_broker_port"`
	MqttBrokerHost string `json:"mqtt_broker_host"`
	MqttUsername   string `json:"mqtt_username"`
	MqttPassword   string `json:"mqtt_password"`
	AliveInterval  int    `json:"alive_interval"`
	InternalLedPin int    `json:"internal_led_pin"`
	CreatedAt      string `json:"created_at"`
}
