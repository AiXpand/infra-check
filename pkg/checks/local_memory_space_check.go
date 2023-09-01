package checks

import (
	"fmt"
	"runtime"
)

type LocalMemorySpaceCheck struct {
	Label     string
	Threshold float64
}

// NewLocalMemorySpaceCheck creates a new LocalMemorySpaceCheck
func NewLocalMemorySpaceCheck(label string, threshold float64) *LocalMemorySpaceCheck {
	return &LocalMemorySpaceCheck{
		Label:     label,
		Threshold: threshold,
	}
}

// Run runs the LocalMemorySpaceCheck
func (d LocalMemorySpaceCheck) Run() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	totalRAMBytes := memStats.Sys
	usedRAMBytes := totalRAMBytes - memStats.HeapIdle
	availableRAMPercentage := (float64(usedRAMBytes) / float64(totalRAMBytes)) * 100

	if availableRAMPercentage > d.Threshold {
		return fmt.Errorf("available RAM is less than %.2f%% (currently %.2f%%)", d.Threshold, availableRAMPercentage)
	}

	return nil
}

// GetLabel returns the label of the LocalMemorySpaceCheck
func (d LocalMemorySpaceCheck) GetLabel() string {
	return d.Label
}
