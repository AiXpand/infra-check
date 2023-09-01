package checks

import (
	"fmt"
	"syscall"
)

type LocalDiskSpaceCheck struct {
	Label              string
	Partition          string
	MaxUsagePercentage float64
}

// NewLocalDiskSpaceCheck creates a new LocalDiskSpaceCheck
func NewLocalDiskSpaceCheck(label, partition string, maxUsagePercentage float64) *LocalDiskSpaceCheck {
	return &LocalDiskSpaceCheck{
		Label:              label,
		Partition:          partition,
		MaxUsagePercentage: maxUsagePercentage,
	}
}

// Run runs the LocalDiskSpaceCheck
func (d LocalDiskSpaceCheck) Run() error {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(d.Partition, &stat)
	if err != nil {
		return fmt.Errorf("error getting disk usage: %v", err)
	}

	// Calculate disk usage in percentage
	usedSpace := float64(stat.Blocks-stat.Bfree) * float64(stat.Bsize)
	totalSpace := float64(stat.Blocks) * float64(stat.Bsize)
	usedPercentage := (usedSpace / totalSpace) * 100

	if usedPercentage > d.MaxUsagePercentage {
		return fmt.Errorf("disk usage exceeds %.2f%% (currently %.2f%%)", d.MaxUsagePercentage, usedPercentage)
	}
	return nil
}

// GetLabel returns the label of the LocalDiskSpaceCheck
func (d LocalDiskSpaceCheck) GetLabel() string {
	return d.Label
}
