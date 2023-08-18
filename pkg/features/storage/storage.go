package storage

import "github.com/aixpand/infra-check/pkg/config"

type StorageClient struct {
	Engine StorageEngine
}

// NewStorageClient creates a new StorageClient
func NewStorageClient(config *config.Config, engine string) *StorageClient {
	var storageEngine StorageEngine
	switch engine {
	case "minio":
		storageEngine = NewMinioClient(config)
		break
	case "local":
		storageEngine = NewLocalStorage()
		break
	}

	return &StorageClient{
		Engine: storageEngine,
	}
}

// FileExists checks if a file exists
func (s *StorageClient) FileExists(path string) (bool, error) {
	return s.Engine.FileExists(path)
}
