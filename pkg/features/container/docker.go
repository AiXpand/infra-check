package container

import (
	"fmt"
	"os/exec"
	"strings"
)

type Docker struct{}

// NewDocker instantiate a new Docker struct
func NewDocker() *Docker {
	return &Docker{}
}

// ContainerExists checks if the container exists
func (d *Docker) ContainerExists(containerName string) error {
	// Run the command to check if the container exists
	cmd := exec.Command("docker", "inspect", "--format={{.Name}}", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	containerExists := strings.TrimSpace(string(output)) != ""
	if !containerExists {
		return fmt.Errorf("container %s does not exist", containerName)
	}

	return nil
}

// ContainerRunning checks if the container is running
func (d *Docker) ContainerRunning(containerName string) error {
	// Run the command to check if the container is running
	cmd := exec.Command("docker", "inspect", "--format={{.State.Running}}", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	containerRunning := strings.TrimSpace(string(output)) == "true"
	if !containerRunning {
		return fmt.Errorf("container %s is not running", containerName)
	}

	return nil
}
