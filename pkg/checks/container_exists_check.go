package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/container"
)

type ContainerExistsCheck struct {
	ContainerName string
	Namespace     string
	Label         string
	Config        *config.Config
}

// NewContainerExistsCheck instantiate a new ContainerExistsCheck struct
func NewContainerExistsCheck(containerName, label, namespace string, config *config.Config) ContainerExistsCheck {
	return ContainerExistsCheck{
		ContainerName: containerName,
		Namespace:     namespace,
		Label:         label,
		Config:        config,
	}
}

// Run executes the check
func (c ContainerExistsCheck) Run() error {
	if err := container.NewContainer(c.Config, c.Namespace).ContainerExists(c.ContainerName); err != nil {
		return err
	}
	return nil
}

// GetLabel returns the label of the check
func (c ContainerExistsCheck) GetLabel() string {
	return c.Label
}
