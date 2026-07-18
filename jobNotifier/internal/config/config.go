package config

import (
	"fmt"
	"jobNotifier/internal/model"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config model.Config //парсинг yaml
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
