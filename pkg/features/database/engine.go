package database

type Engine interface {
	Connect(username, password, database string) error
	Disconnect()
	Ping() error
}
