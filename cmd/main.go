package main

import (
	"fmt"
	"github.com/aixpand/infra-check/pkg/checks"
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/pterm/pterm"
	"os"
)

var AppVersion = "1.0.0"

func main() {
	cfg, err := config.NewConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print out a header
	pterm.DefaultHeader.WithFullWidth().Println("AiExpand Infrastructure Checker " + AppVersion)
	pterm.Println() // spacer

	var currentChecks []checks.Check
	currentChecks = initialiseChecks(cfg, currentChecks)

	// Create progressbar as fork from the default progressbar.
	p, _ := pterm.DefaultProgressbar.WithTotal(len(currentChecks)).WithTitle("Checking infrastructure").Start()
	for _, c := range currentChecks {
		p.UpdateTitle("Running check " + c.GetLabel())
		if err = c.Run(); err != nil {
			pterm.Error.Println(c.GetLabel() + " failed with error: " + err.Error())
			p.Increment()
			continue
		}
		pterm.Success.Println(c.GetLabel())
		p.Increment()
	}

}

// initialiseChecks creates the checks based on the configuration
func initialiseChecks(cfg *config.Config, currentChecks []checks.Check) []checks.Check {
	for _, chk := range cfg.Checks {
		switch chk.Type {
		case "dummy_check":
			currentChecks = append(currentChecks, checks.DummyCheck{Label: chk.Label})
			break
		case "container_exists_check":
			currentChecks = append(currentChecks, checks.NewContainerExistsCheck(chk.ContainerName, chk.Label, chk.Namespace, cfg))
			break
		case "container_running_check":
			currentChecks = append(currentChecks, checks.NewContainerRunningCheck(chk.ContainerName, chk.Label, chk.Namespace, cfg))
			break
		case "http_response_check":
			currentChecks = append(currentChecks, checks.HttpResponseCheck{Url: chk.Url, Label: chk.Label, ExpectedResponse: chk.Code})
			break
		case "mqtt_connection_check":
			currentChecks = append(currentChecks, checks.NewMqttConnectionCheck(cfg, chk.Label))
			break
		case "redis_connection_check":
			currentChecks = append(currentChecks, checks.NewRedisConnectionCheck(cfg, chk.Label))
			break
		case "database_connection_check":
			currentChecks = append(currentChecks, checks.NewDatabaseConnectionCheck(cfg, chk.Label, chk.Username, chk.Password, chk.Database))
			break
		case "file_exists_check":
			currentChecks = append(currentChecks, checks.NewFileExistsCheck(cfg, chk.Label, chk.Path, chk.Engine))
			break
		case "execution_engine_heartbeat_check":
			currentChecks = append(currentChecks, checks.NewExecutionEngineHeartbeatCheck(chk.BoxName, chk.Label, cfg))
			break
		case "terminus_check":
			currentChecks = append(currentChecks, checks.NewTerminusCheck(chk.Label, chk.Url))
			break
		case "local_disk_space_check":
			currentChecks = append(currentChecks, checks.NewLocalDiskSpaceCheck(chk.Label, chk.Path, chk.Threshold))
			break
		case "local_memory_space_check":
			currentChecks = append(currentChecks, checks.NewLocalMemorySpaceCheck(chk.Label, chk.Threshold))
			break
		}
	}
	return currentChecks
}
