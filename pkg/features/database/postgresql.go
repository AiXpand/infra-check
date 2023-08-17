package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aixpand/infra-check/pkg/config"
	_ "github.com/lib/pq"
	"time"
)

type Postgresql struct {
	DatabaseClient *sql.DB
	Hostname       string
	Port           int
}

// NewPostgresql creates a new Postgresql
func NewPostgresql(config *config.Config) *Postgresql {
	return &Postgresql{
		DatabaseClient: nil,
		Hostname:       config.Database.Host,
		Port:           config.Database.Port,
	}
}

func (p *Postgresql) Connect(username, password, database string) error {
	// build database connection string
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, p.Hostname, p.Port, database)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	p.DatabaseClient = db
	return nil
}

func (p *Postgresql) Disconnect() {
	if p.DatabaseClient != nil {
		_ = p.DatabaseClient.Close()
	}
}

func (p *Postgresql) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result string
	err := p.DatabaseClient.QueryRowContext(ctx, "SELECT 'Connected'").Scan(&result)
	if err != nil {
		return err
	}

	return nil
}
