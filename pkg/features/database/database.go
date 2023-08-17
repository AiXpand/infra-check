package database

import "github.com/aixpand/infra-check/pkg/config"

type DatabaseClient struct {
	Engine Engine
}

// NewDatabaseClient creates a new DatabaseClient
func NewDatabaseClient(config *config.Config) *DatabaseClient {
	var databaseEngine Engine
	switch config.Database.Engine {
	case "postgresql":
		databaseEngine = NewPostgresql(config)
		break
	}

	return &DatabaseClient{
		Engine: databaseEngine,
	}
}

// Connect connects to the database
func (d *DatabaseClient) Connect(username, password, database string) error {
	return d.Engine.Connect(username, password, database)
}

// Disconnect disconnects from the database
func (d *DatabaseClient) Disconnect() {
	d.Engine.Disconnect()
}

// Ping pings the database
func (d *DatabaseClient) Ping() error {
	return d.Engine.Ping()
}
