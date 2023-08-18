package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	Hostname string
	Username string
	Password string
	Port     int
	Client   MQTT.Client
}

// NewMqttClient creates a new MQTT client
func NewMqttClient(hostname string, username string, password string, port int) *MqttClient {
	return &MqttClient{
		Hostname: hostname,
		Username: username,
		Password: password,
		Port:     port,
	}
}

// Connect connects to the MQTT server
func (m *MqttClient) Connect() error {
	// Create the connection string
	serverURI := fmt.Sprintf("tcp://%s:%d", m.Hostname, m.Port)

	// Create an MQTT client options
	opts := MQTT.NewClientOptions()
	opts.AddBroker(serverURI)
	opts.SetUsername(m.Username)
	opts.SetPassword(m.Password)

	// Create an MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	m.Client = client
	return nil
}

// Disconnect disconnects from the MQTT server
func (m *MqttClient) Disconnect() {
	m.Client.Disconnect(250)
}

func (m *MqttClient) Subscribe(topic string, callback MQTT.MessageHandler) error {
	if token := m.Client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *MqttClient) Unsubscribe(topic string) error {
	if token := m.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
