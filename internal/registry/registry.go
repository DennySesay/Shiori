package registry

import (
	"fmt"
	"log"
	"os"
)

func FindConfig() {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		log.Fatal(dirErr)
	} else {
		fmt.Printf("config directory %v", dir)
	}
}
