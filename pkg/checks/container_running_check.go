package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/container"
)

type ContainerRunningCheck struct {
	ContainerName string
	Namespace     string
	Label         string
	Config        *config.Config
}

// NewContainerRunningCheck instantiate a new ContainerRunningCheck struct
func NewContainerRunningCheck(containerName, label, namespace string, config *config.Config) *ContainerRunningCheck {
	return &ContainerRunningCheck{
		ContainerName: containerName,
		Namespace:     namespace,
		Label:         label,
		Config:        config,
	}
}

// Run the check
func (c *ContainerRunningCheck) Run() error {
	err := container.NewContainer(c.Config, c.Namespace).ContainerRunning(c.ContainerName)
	if err != nil {
		return err
	}

	return nil
}

// GetLabel returns the label of the check
func (c *ContainerRunningCheck) GetLabel() string {
	return c.Label
}
