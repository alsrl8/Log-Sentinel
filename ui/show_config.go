package ui

import (
	"LogSentinel/config"
	"LogSentinel/utils"
	"fmt"
)

func showConfig() {
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.GetLogger().Error("Failed to load config: %v", err)
		return
	}

	fmt.Println("Current Configuration")
	fmt.Println("====================")
	for i, kube := range cfg.Kube {
		fmt.Printf("%d. Service: %s (%s)\n", i+1, kube.Name, kube.ServiceLabel)
	}
}
