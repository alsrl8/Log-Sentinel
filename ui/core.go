package ui

import (
	"LogSentinel/fetch"
	"LogSentinel/utils"
	"fmt"
	"github.com/manifoldco/promptui"
)

func ShowMenu() {
	logger := utils.GetLogger()

	for {
		clearScreen()
		fmt.Println("LogSentinel - Kubernetes Log Monitor")
		fmt.Println("===================================")

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				"Fetch Pod Logs",
				"List Available Pods",
				"Show Configuration",
				"Exit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			logger.Error("Prompt failed: %v", err)
			return
		}

		clearScreen()
		switch result {
		case "Fetch Pod Logs":
			fetch.K8sPodLogs()
		case "List Available Pods":
			listPods()
		case "Show Configuration":
			showConfig()
		case "Exit":
			logger.Info("Exiting LogSentinel")
			return
		}

		// Pause to show results before clearing
		fmt.Printf("\nPress Enter to continue...")
		_, err = fmt.Scanln()
		if err != nil {
			logger.Error("Failed to read input: %v", err)
			return
		}
	}
}
