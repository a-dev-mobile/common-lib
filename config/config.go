package config

import (
    "errors"
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

// GetConfig loads configuration from file
func GetConfig[T any](configPath string, configName string) (*T, error) {
    configFile := filepath.Join(configPath, configName)

    if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
        return nil, fmt.Errorf("config file does not exist: %s", configFile)
    }

    return loadConfig[T](configFile)
}

// loadConfig reads and decodes the YAML configuration file.
func loadConfig[T any](file string) (*T, error) {
    configFile, err := os.ReadFile(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    var config T
    if err := yaml.Unmarshal(configFile, &config); err != nil {
        return nil, fmt.Errorf("error decoding config file: %w", err)
    }

    return &config, nil
}
