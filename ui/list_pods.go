package ui

import (
	"LogSentinel/config"
	"LogSentinel/utils"
	"fmt"
)

func listPods() {
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.GetLogger().Error("Failed to load config: %v", err)
		return
	}

	fmt.Println("Configured Pods")
	fmt.Println("===============")
	for _, kube := range cfg.Kube {
		fmt.Printf("\nName: %s\nNamespace: %s\nService Label: %s\nLog Path: %s\nFormat: %s\n",
			kube.Name, kube.NameSpace, kube.ServiceLabel, kube.LogPath, kube.Format)
		fmt.Println("---------------")
	}
}
