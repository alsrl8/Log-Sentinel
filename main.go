package main

import (
	"LogSentinel/config"
	"log"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config\n%v", err)
	}

	//watcher := collector.LogWatcher{
	//	FilePath: "./temp.log",
	//	Offset:   "",
	//}
	//if err := watcher.Watch(); err != nil {
	//	panic(err)
	//}

}
