package storage

// StorageEngine is the interface for the storage feature
type StorageEngine interface {
	// FileExists checks if a file exists
	FileExists(path string) (bool, error)
}
