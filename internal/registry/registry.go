package registry

import (
	"fmt"
	"log"
	"os"
	path "path/filepath"
)

func GetRegistryPath() (string, error) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	var registryPath string = path.Join(dir, "shiori", "registry.json")
	fmt.Println(registryPath)
	return registryPath, dirErr
}

func LoadRegistry() {}
func SaveRegistry() {}
