package collector

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

type Watcher interface {
	Watch() error
}

type LogWatcher struct {
	FilePath string
	Offset   string
}

func (w *LogWatcher) Watch() error {
	if w.FilePath == "" {
		return errors.New("file path is empty")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer func(watcher *fsnotify.Watcher) {
		_ = watcher.Close()
	}(watcher)

	file, err := os.Open(w.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	err = watcher.Add(w.FilePath)
	if err != nil {
		return fmt.Errorf("failed to add file to watcher: %w", err)
	}

	fmt.Printf("Watching file %s\n", w.FilePath)
	//reader := bufio.NewReader(file)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File %s was modified\n", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil

			}
			fmt.Println("error:", err)
		}
	}
}
