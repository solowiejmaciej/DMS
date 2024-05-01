package services

import (
	"dms/repositories"
	"sync"
	"time"
)

var LastMessageTimes = make(map[string]time.Time)
var Mutex = &sync.Mutex{}

func UpdateLastMessageTime(deviceId string) {
	Mutex.Lock()
	LastMessageTimes[deviceId] = time.Now()
	Mutex.Unlock()
}

func CalculateDeviceStatus(deviceId string) string {
	Mutex.Lock()
	lastMessageTime, ok := LastMessageTimes[deviceId]
	Mutex.Unlock()

	if !ok {
		return "OFFLINE"
	}

	deviceConfiguration, err := repositories.GetConfigurationByDeviceId(deviceId)
	if err != nil {
		return "OFFLINE"
	}

	if time.Since(lastMessageTime).Minutes() <= float64(deviceConfiguration.AliveInterval) {
		return "ONLINE"
	} else {
		return "OFFLINE"
	}
}
