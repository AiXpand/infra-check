package container

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Kubernetes struct {
	Namespace string
}

type ContainerStatus struct {
	State struct {
		Running struct {
			StartedAt string `json:"startedAt"`
		} `json:"running"`
	} `json:"state"`
}

type PodStatus struct {
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
}

// NewKubernetes instantiate a new Kubernetes struct
func NewKubernetes(namespace string) *Kubernetes {
	return &Kubernetes{}
}

// ContainerExists checks if the container exists in Kubernetes
func (k *Kubernetes) ContainerExists(containerName string) error {
	// Run the kubectl command to check if the container exists
	cmdArgs := []string{"get", "pod", "--namespace", k.Namespace, "-o", "json"}
	cmd := exec.Command("kubectl", cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	if strings.Contains(string(output), containerName) {
		return nil // Container exists
	}

	return fmt.Errorf("container %s does not exist in namespace %s", containerName, k.Namespace)
}

// ContainerRunning retrieves the container status from Kubernetes
func (k *Kubernetes) ContainerRunning(containerName string) error {
	// Run the kubectl command to get the pod details in JSON format
	cmdArgs := []string{"get", "pod", "--namespace", k.Namespace, "-o", "json"}
	cmd := exec.Command("kubectl", cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Parse the pod status from the JSON output
	podStatus, parseErr := parsePodStatus(output)
	if parseErr != nil {
		return parseErr
	}

	// Check if the container status is "Running"
	for _, containerStatus := range podStatus.ContainerStatuses {
		if containerStatus.State.Running.StartedAt != "" {
			return nil
		}
	}

	return fmt.Errorf("pod: %s is not running", containerName)
}

// parsePodStatus parses the pod status from the JSON output
func parsePodStatus(jsonData []byte) (PodStatus, error) {
	var podStatus PodStatus
	err := json.Unmarshal(jsonData, &podStatus)
	if err != nil {
		return PodStatus{}, fmt.Errorf("error parsing JSON: %v", err)
	}
	return podStatus, nil
}
