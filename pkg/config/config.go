package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Engine string `json:"engine"`
	Redis  struct {
		Host     string `json:"host"`
		Password string `json:"password,omitempty"`
		Port     int    `json:"port"`
	}
	Database struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Engine string `json:"engine"`
	}
	Mqtt struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Minio struct {
		Host      string `json:"host"`
		Port      int    `json:"port"`
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
		UseSSL    bool   `json:"useSSL"`
	}
	Checks []struct {
		Type          string `json:"type"`
		Label         string `json:"label"`
		ContainerName string `json:"container_name,omitempty"`
		Url           string `json:"url,omitempty"`
		Code          int    `json:"code,omitempty"`
		Username      string `json:"username,omitempty"`
		Password      string `json:"password,omitempty"`
		Database      string `json:"database,omitempty"`
		Path          string `json:"path,omitempty"`
		Engine        string `json:"engine,omitempty"`
		BoxName       string `json:"box_name,omitempty"`
		Namespace     string `json:"namespace,omitempty"`
	} `json:"checks"`
}

// NewConfiguration loads configuration form file
func NewConfiguration() (*Config, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configuration Config
	err = json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
