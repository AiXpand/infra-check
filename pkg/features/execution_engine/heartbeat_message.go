package execution_engine

type HeartbeatMessage struct {
	Type          string   `json:"type"`
	EEPayloadPath []string `json:"EE_PAYLOAD_PATH"`
}
