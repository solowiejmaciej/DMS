package initializers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var Client mqtt.Client

func ConfigureMqtt() {
	log.Info("Connecting to MQTT broker")
	var broker = "d7f63eb1c4a14248a3dfc2d032fbb00c.s1.eu.hivemq.cloud" // find the host name in the Overview of your cluster (see readme)
	var port = 8883                                                    // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("DMS")   // set a name as you desire
	opts.SetUsername("esp32") // these are the credentials that you declare for your cluster
	opts.SetPassword("zaq1@WSX")

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker:", token.Error()))
	}
	log.Info("Connection to MQTT broker established")
}
