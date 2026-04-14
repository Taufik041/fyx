package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds everything fyx needs to know about the user's setup.
type Config struct {
	Provider string `json:"provider"`
	APIKey   string `json:"api_key"`
	Active   bool   `json:"active"`
	Theme    string `json:"theme"`
}

// configDir returns the path to ~/.fyx/
func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".fyx"), nil
}

// configPath returns the path to ~/.fyx/config.json
func configPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

// Load reads the config file and returns a Config struct.
// If the file doesn't exist yet, it returns a default empty Config.
func Load() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// First time — return empty config, not an error
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Save writes the Config struct to ~/.fyx/config.json.
// It creates the ~/.fyx/ directory if it doesn't exist yet.
func Save(cfg *Config) error {
	dir, err := configDir()
	if err != nil {
		return err
	}

	// Create ~/.fyx/ if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	// 0600 means only the current user can read/write — important since it stores API keys
	return os.WriteFile(path, data, 0600)
}
