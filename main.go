package main

import (
	"LogSentinel/config"
	"LogSentinel/ui"
	"LogSentinel/utils"
	"sync"
)

func main() {
	logger := utils.GetLogger()
	logger.Info("LogSentinel started")

	_, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ui.ShowMenu()
		wg.Done()
	}()
	wg.Wait()
	logger.Info("LogSentinel stopped")
}
