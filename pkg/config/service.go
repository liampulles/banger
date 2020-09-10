package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/liampulles/banger/pkg/file"
)

type ConfigService interface {
	Load() (Config, error)
	Save(Config) error
}

type ConfigServiceImpl struct {
	configPath string
}

var _ ConfigService = &ConfigServiceImpl{}

func NewConfigServiceImpl(configPath string) *ConfigServiceImpl {
	return &ConfigServiceImpl{
		configPath: configPath,
	}
}

func (cs *ConfigServiceImpl) Load() (Config, error) {
	bytes, err := file.ReadBytes(cs.configPath)
	if err != nil {
		return nil, fmt.Errorf("could not load config - read error: %w", err)
	}

	var cfg ConfigImpl
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("could not read config - unmarshal error: %w", err)
	}

	return &cfg, nil
}

func (cs *ConfigServiceImpl) Save(cfg Config) error {
	bytes, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return fmt.Errorf("could not write config - marshal error: %w", err)
	}
	fmt.Println(cfg)

	if err := file.WriteBytes(cs.configPath, bytes); err != nil {
		return fmt.Errorf("could not write config - writeBytes error: %w", err)
	}
	return nil
}

func DefaultConfigPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("could not get default config path - user config error: %w", err)
	}

	return path.Join(userConfigDir, "banger", "config"), nil
}
