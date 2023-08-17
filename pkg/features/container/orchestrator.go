package container

type Orchestrator interface {
	ContainerExists(containerName string) error
	ContainerRunning(containerName string) error
}
