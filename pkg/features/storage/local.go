package storage

import "os"

// LocalStorage is the interface for the storage engine
type LocalStorage struct{}

// NewLocalStorage creates a new StorageEngine
func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

// FileExists checks if a file exists
func (s *LocalStorage) FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return true, nil
}
