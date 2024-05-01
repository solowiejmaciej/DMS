package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"notificator/events"
	"notificator/repositories"
	"os"
	"strings"
)

func mapCodeNotifications(body []byte) events.ConfirmPhoneNumber {
	var event events.ConfirmPhoneNumber
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Errorf("Error unmarshalling notification: %v", err)
	}
	return event
}

func ProcessCodeNotifications(message []byte) {
	event := mapCodeNotifications(message)
	log.Info("Processing event: ", event.EventId)
	text := "Your DMS code: " + event.Code
	processNotification(event.UserId, text)
}

func processNotification(userId uint, message string) {
	log.Infof("Sending notification to user %v", userId)
	number, err := repositories.GetUserPhoneNumber(userId)
	if err != nil {
		log.Errorf("Error fetching user data: %v", err)
		return
	}
	log.Infof("Sending SMS to %v", number)
	sendSms(number, message)
}

func sendSms(phoneNumber string, message string) {
	log.Infof("Sending SMS to %v", phoneNumber)
	data := url.Values{}
	data.Set("key", os.Getenv("SMS_KEY"))
	data.Set("password", os.Getenv("SMS_PASSWORD"))
	data.Set("from", os.Getenv("SMS_SENDER"))
	data.Set("to", phoneNumber)
	data.Set("msg", message)

	client := &http.Client{}
	apiUrl := os.Getenv("SMS_API_URL") + "/sms"
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Errorf("Error creating new request: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error while making the request %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error while sending SMS: %v", resp.Status)
	}
}
