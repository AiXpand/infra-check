package checks

type DummyCheck struct {
	Label string
}

// NewDummyCheck creates a new DummyCheck
func NewDummyCheck(label string) *DummyCheck {
	return &DummyCheck{
		Label: label,
	}
}

// Run runs the DummyCheck
func (d DummyCheck) Run() error {
	return nil
}

// GetLabel returns the label of the DummyCheck
func (d DummyCheck) GetLabel() string {
	return d.Label
}
