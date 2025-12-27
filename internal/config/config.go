package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL 	        string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	full_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Home directory not found: %w", err)
	}

	data, err := os.ReadFile(full_path)
	if err != nil {
		return Config{}, fmt.Errorf("Could not read file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("Could not unmarshal JSON: %w", err)
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"

	home_path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home_path + "/" + configFileName, nil
}

func write(cfg *Config) error {

	full_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Home directory not found: %w", err)
	}

	jsonData, err := json.Marshal(*cfg)
	if err != nil {
		return fmt.Errorf("Cannot marshal data: %w", err)
	}

	if err := os.WriteFile(full_path, jsonData, 0644); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}

	return nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	if err := write(c); err != nil {
		return fmt.Errorf("Could not write to JSON: %w", err)
	}
	return nil
}
