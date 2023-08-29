package checks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Terminus struct {
	Status  string      `json:"status"`
	Info    interface{} `json:"info"`
	Error   interface{} `json:"error"`
	Details interface{} `json:"details"`
}

type TerminusCheck struct {
	Label string
	Url   string
}

func NewTerminusCheck(label, url string) *TerminusCheck {
	return &TerminusCheck{
		Label: label,
		Url:   url,
	}
}

// Run executes the check
func (c TerminusCheck) Run() error {
	resp, err := http.Get(c.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var terminus Terminus
		err = json.Unmarshal(body, &terminus)
		if err != nil {
			return err
		}

		// Handling Error field if it is a map with dynamic keys
		if errMap, ok := terminus.Error.(map[string]interface{}); ok {
			for key, value := range errMap {
				if nestedMap, nestedMapOk := value.(map[string]interface{}); nestedMapOk {
					status, statusOk := nestedMap["status"].(string)
					if statusOk {
						return fmt.Errorf("service: %s, status: %s\n", key, status)
					}
				}
			}
		}
	}
	return nil
}

// GetLabel returns the label of the check
func (c TerminusCheck) GetLabel() string {
	return c.Label
}
