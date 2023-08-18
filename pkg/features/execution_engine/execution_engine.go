package execution_engine

import (
	"encoding/json"
	"fmt"
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/mqtt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type ExecutionEngine struct {
	Config          *config.Config
	Client          *mqtt.MqttClient
	messageReceived chan bool
	BoxName         string
}

// NewExecutionEngine creates a new ExecutionEngine
func NewExecutionEngine(config *config.Config, boxName string) *ExecutionEngine {
	client := mqtt.NewMqttClient(config.Mqtt.Host, config.Mqtt.Username, config.Mqtt.Password, config.Mqtt.Port)
	e := &ExecutionEngine{
		Config:          config,
		Client:          client,
		messageReceived: make(chan bool),
		BoxName:         boxName,
	}

	return e
}

func (e *ExecutionEngine) onMessageReceived(client paho.Client, message paho.Message) {
	var payload HeartbeatMessage
	err := json.Unmarshal(message.Payload(), &payload)
	if err != nil {
		log.Println(err)
		return
	}

	if payload.EEPayloadPath[0] == e.BoxName && payload.Type == "heartbeat" {
		e.messageReceived <- true
	}
}

func (e *ExecutionEngine) WaitForHeartbeat() error {
	waitTimeout := 60 * time.Second

	err := e.Client.Connect()
	if err != nil {
		return err
	}

	// Subscribe to the heartbeat topic
	if err = e.Client.Subscribe("lummetry/ctrl", e.onMessageReceived); err != nil {
		return err
	}

	// Wait for the heartbeat
	select {
	case <-e.messageReceived:
		_ = e.Client.Unsubscribe("lummetry/ctrl")
		e.Client.Disconnect()
		return nil
	case <-time.After(waitTimeout):
		_ = e.Client.Unsubscribe("lummetry/ctrl")
		e.Client.Disconnect()
		return fmt.Errorf("timeout waiting for expected message")
	}
}
