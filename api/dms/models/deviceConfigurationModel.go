package models

type DeviceConfigurationModel struct {
	MqttBrokerPort int    `json:"mqqt_broker_port"`
	MqttBrokerHost string `json:"mqqt_broker_host"`
	MqttUsername   string `json:"mqqt_username"`
	MqttPassword   string `json:"mqqt_password"`
	AliveInterval  int    `json:"alive_interval"`
	InternalLedPin int    `json:"internal_led_pin"`
}
