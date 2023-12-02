package config

import (
    "errors"
    // Errors package for handling errors.
    "fmt"
    // Fmt package for formatting strings.
    "os"
    // Os package for interacting with the operating system, like file handling.
    "path/filepath"
    // Filepath package for manipulating filename paths.
    "gopkg.in/yaml.v3"
    // Yaml.v3 package for YAML processing.
)

// GetConfig loads configuration from file.
// This function is generic and works with any type T.
// It takes a configuration path and name, then returns a pointer to the config struct or an error.
func GetConfig[T any](configPath string, configName string) (*T, error) {
    configFile := filepath.Join(configPath, configName)
    // Joins the path and filename to create the full path to the configuration file.

    if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
        // Checks if the file exists. If it does not, returns an error.
        return nil, fmt.Errorf("config file does not exist: %s", configFile)
    }

    // If the file exists, calls loadConfig to read and parse the file.
    return loadConfig[T](configFile)
}

// loadConfig reads and decodes the YAML configuration file.
// It is a private function, indicated by the lowercase first letter.
// Takes the file path as input and returns a pointer to the config struct or an error.
func loadConfig[T any](file string) (*T, error) {
    configFile, err := os.ReadFile(file)
    // Reads the file. If there is an error reading, it returns an error.
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    var config T
    // Declares a variable of type T to hold the configuration data.

    if err := yaml.Unmarshal(configFile, &config); err != nil {
        // Tries to unmarshal the YAML file into the config variable. If it fails, returns an error.
        return nil, fmt.Errorf("error decoding config file: %w", err)
    }

    // Returns a pointer to the config struct if successful.
    return &config, nil
}
