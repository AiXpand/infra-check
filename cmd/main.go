package main

import (
	"fmt"
	"github.com/aixpand/infra-check/pkg/checks"
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/pterm/pterm"
	"os"
)

const VERSION = "1.0.0"

func main() {
	cfg, err := config.NewConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// build a version string from the environment
	version := os.Getenv("APP_VERSION")
	if version != "" {
		version = VERSION
	}

	// print out a header
	pterm.DefaultHeader.WithFullWidth().Println("CAVI Infrastructure Checker " + version)
	pterm.Println() // spacer

	var currentChecks []checks.Check
	for _, chk := range cfg.Checks {
		switch chk.Type {
		case "dummy_check":
			currentChecks = append(currentChecks, checks.DummyCheck{Label: chk.Label})
			break
		case "container_exists_check":
			currentChecks = append(currentChecks, checks.NewContainerExistsCheck(chk.ContainerName, chk.Label, cfg))
			break
		case "container_running_check":
			currentChecks = append(currentChecks, checks.NewContainerRunningCheck(chk.ContainerName, chk.Label, cfg))
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
		}
	}

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
