package registry

import (
	"os"
	path "path/filepath"
)

type Registry struct {
	Entry map[string]Entry
}

func GetRegistryPath() (string, error) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		return "", dirErr
	}

	var registryPath string = path.Join(dir, "shiori", "registry.json")
	return registryPath, dirErr
}

// return a pointer of register
func LoadRegistry() (*Registry, error) {
	// call get reg then read file then parse contents then assign to reg struct and return
	registryPath, err := GetRegistryPath()
	if err != nil {
		return &Registry{}, err
	}

	registry, err := os.Open(registryPath)
	if err != nil {
		return &Registry{}, err
	}
	defer registry.Close()
}

func SaveRegistry() {}
