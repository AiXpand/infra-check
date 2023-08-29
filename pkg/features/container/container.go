package container

import "github.com/aixpand/infra-check/pkg/config"

type Container struct {
	Engine       string
	Orchestrator Orchestrator
}

// NewContainer instantiate a new Container struct
func NewContainer(config *config.Config, namespace string) *Container {
	var orchestrator Orchestrator
	switch config.Engine {
	case "podman":
		orchestrator = NewPodman()
		break
	case "docker":
		orchestrator = NewDocker()
		break
	case "kubernetes":
		orchestrator = NewKubernetes(namespace)
		break
	}

	return &Container{
		Engine:       config.Engine,
		Orchestrator: orchestrator,
	}
}

func (c *Container) ContainerExists(containerName string) error {
	if err := c.Orchestrator.ContainerExists(containerName); err != nil {
		return err
	}
	return nil
}

func (c *Container) ContainerRunning(containerName string) error {
	if err := c.Orchestrator.ContainerRunning(containerName); err != nil {
		return err
	}
	return nil
}
