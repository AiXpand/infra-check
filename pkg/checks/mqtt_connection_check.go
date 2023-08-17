package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/mqtt"
)

type MqttConnectionCheck struct {
	Hostname string
	Username string
	Password string
	Port     int
	Label    string
}

// NewMqttConnectionCheck instantiate a new MqttConnectionCheck struct
func NewMqttConnectionCheck(config *config.Config, label string) *MqttConnectionCheck {
	return &MqttConnectionCheck{
		Hostname: config.Mqtt.Host,
		Username: config.Mqtt.Username,
		Password: config.Mqtt.Password,
		Port:     config.Mqtt.Port,
		Label:    label,
	}
}

// Run executes the check
func (m MqttConnectionCheck) Run() error {
	client := mqtt.NewMqttClient(m.Hostname, m.Username, m.Password, m.Port)
	if err := client.Connect(); err != nil {
		return err
	}

	client.Disconnect()
	return nil
}

// GetLabel returns the label of the check
func (m MqttConnectionCheck) GetLabel() string {
	return m.Label
}
