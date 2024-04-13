package initializers

import (
	"dms/repositories"
	"dms/services"
)

func ConfigureMqtt() {
	configurations, err := repositories.GetAllConfigurations()
	if err != nil {
		return
	}
	for _, configuration := range configurations {
		services.ConfigureMqttForDevice(configuration)
	}
}
