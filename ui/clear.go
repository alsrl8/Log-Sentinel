package ui

import (
	"LogSentinel/utils"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	logger := utils.GetLogger()
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			logger.Error(err.Error())
			return
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			logger.Error(err.Error())
			return
		}
	default:
		logger.Warn("Unsupported OS")
		return
	}

}
