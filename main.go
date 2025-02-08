package main

import (
	"fmt"

	"github.com/modestprophet/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config: %v", err)
		return
	}

	err = cfg.SetUser("modest")
	if err != nil {
		fmt.Printf("Failed to update config: %v", err)
		return
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to re-read config: %v", err)
	}

	fmt.Printf("Final configuration:\n%+v\n", updatedCfg)
}
