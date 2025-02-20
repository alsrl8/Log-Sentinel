package main

import (
	"LogSentinel/config"
	"LogSentinel/fetch"
	"LogSentinel/utils"
)

func main() {
	logger := utils.GetLogger()
	logger.Info("LogSentinel started")

	_, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	fetch.K8sPodLogs()
}
