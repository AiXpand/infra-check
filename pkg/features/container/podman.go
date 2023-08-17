package container

import (
	"fmt"
	"os/exec"
	"strings"
)

type Podman struct{}

// NewPodman instantiate a new Podman struct
func NewPodman() *Podman {
	return &Podman{}
}

// ContainerExists checks if the container exists
func (pm *Podman) ContainerExists(containerName string) error {
	// Run the command to check if the container exists
	cmd := exec.Command("podman", "inspect", "--format={{.Name}}", containerName)
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
func (pm *Podman) ContainerRunning(containerName string) error {
	// Run the command to check if the container is running
	cmd := exec.Command("podman", "inspect", "--format={{.State.Status}}", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	containerRunning := strings.TrimSpace(string(output)) == "running"
	if !containerRunning {
		return fmt.Errorf("container %s is not running", containerName)
	}
	return nil
}
