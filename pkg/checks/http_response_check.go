package checks

import (
	"fmt"
	"net/http"
)

type HttpResponseCheck struct {
	Url              string
	Label            string
	ExpectedResponse int
}

// Run executes the check
func (c HttpResponseCheck) Run() error {
	resp, err := http.Get(c.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != c.ExpectedResponse {
		return fmt.Errorf("expected response code %d, got %d instead", c.ExpectedResponse, resp.StatusCode)
	}
	return nil
}

// GetLabel returns the label of the check
func (c HttpResponseCheck) GetLabel() string {
	return c.Label
}
