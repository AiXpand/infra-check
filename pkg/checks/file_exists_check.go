package checks

import (
	"github.com/aixpand/infra-check/pkg/config"
	"github.com/aixpand/infra-check/pkg/features/storage"
)

type FileExistsCheck struct {
	Config        *config.Config
	Label         string
	Path          string
	Engine        string
	StorageEngine storage.StorageEngine
}

// NewFileExistsCheck creates a new FileExistsCheck
func NewFileExistsCheck(config *config.Config, label string, path string, engine string) *FileExistsCheck {
	storageEngine := storage.NewStorageClient(config, engine)
	return &FileExistsCheck{
		Config:        config,
		Label:         label,
		Path:          path,
		Engine:        engine,
		StorageEngine: storageEngine,
	}
}

// Run executes the check
func (c *FileExistsCheck) Run() error {
	_, err := c.StorageEngine.FileExists(c.Path)
	if err != nil {
		return err
	}

	return nil
}

// GetLabel returns the label of the check
func (c *FileExistsCheck) GetLabel() string {
	return c.Label
}
