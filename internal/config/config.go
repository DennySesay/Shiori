package config

import (
	"fmt"
	"os"
	path "path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type ShioriConfig struct {
	ShioriSettings Settings                 `toml:"settings"`
	Entry          map[string]RegistryEntry `toml:"registry"`
}

type Settings struct {
	AlwaysOpenDir      bool   `toml:"always_open_dir"`
	ConflictResolution string `toml:"conflict_resolution"`
}

type RegistryEntry struct {
	Path   string `toml:"path"`
	Editor string `toml:"editor"`
}

func GetConfig() (string, error) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		return "", dirErr
	}

	var registryPath string = path.Join(dir, "shiori", "config.toml")
	return registryPath, dirErr
}

func LoadConfig() (ShioriConfig, error) {
	registryPath, err := GetConfig()
	if err != nil {
		return ShioriConfig{}, err
	}

	content, err := os.ReadFile(registryPath)

	var config ShioriConfig
	err = toml.Unmarshal(content, &config)
	if err != nil {
		return ShioriConfig{}, err
	}

	fmt.Printf("%+v\n", config)

	return config, nil
}

func SaveRegistry() {}
