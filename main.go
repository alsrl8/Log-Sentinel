package main

import (
	"LogSentinel/config"
	"LogSentinel/fetch"
	"LogSentinel/program_args"
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

	args := program_args.GetArgs()
	if args.DisableUI {
		logger.Info("Running in background, fetching logs")
		fetch.K8sPodLogs()
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
