package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/execution_engine"
)

type ExecutionEngineHeartbeatCheck struct {
	ExecutionEngineName string
	Label               string
	Config              *config.Config
}

// NewExecutionEngineHeartbeatCheck creates a new ExecutionEngineHeartbeatCheck
func NewExecutionEngineHeartbeatCheck(executionEngineName string, label string, config *config.Config) *ExecutionEngineHeartbeatCheck {
	return &ExecutionEngineHeartbeatCheck{
		ExecutionEngineName: executionEngineName,
		Label:               label,
		Config:              config,
	}
}

// Run runs the ExecutionEngineHeartbeatCheck
func (e *ExecutionEngineHeartbeatCheck) Run() error {
	ee := execution_engine.NewExecutionEngine(e.Config, e.ExecutionEngineName)
	if err := ee.WaitForHeartbeat(); err != nil {
		return err
	}

	return nil
}

// GetLabel returns the label of the ExecutionEngineHeartbeatCheck
func (e *ExecutionEngineHeartbeatCheck) GetLabel() string {
	return e.Label
}
