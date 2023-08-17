package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/database"
)

type DatabaseConnectionCheck struct {
	Label    string
	Config   *config.Config
	Username string
	Password string
	Database string
}

// NewDatabaseConnectionCheck creates a new DatabaseConnectionCheck
func NewDatabaseConnectionCheck(config *config.Config, label string, username string, password string, database string) *DatabaseConnectionCheck {
	return &DatabaseConnectionCheck{
		Label:    label,
		Config:   config,
		Username: username,
		Password: password,
		Database: database,
	}
}

// Run runs the DatabaseConnectionCheck
func (d *DatabaseConnectionCheck) Run() error {
	databaseClient := database.NewDatabaseClient(d.Config)
	if err := databaseClient.Connect(d.Username, d.Password, d.Database); err != nil {
		return err
	}
	defer databaseClient.Disconnect()

	if err := databaseClient.Ping(); err != nil {
		return err
	}

	return nil
}

// GetLabel returns the label of the DatabaseConnectionCheck
func (d *DatabaseConnectionCheck) GetLabel() string {
	return d.Label
}
