package fetch

import (
	"LogSentinel/utils"
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type LogComparison struct {
	IsSame     bool
	NeedsReset bool
}

func compareDiff(localPath string, remoteContent []byte) (*LogComparison, error) {
	logger := utils.GetLogger()
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		return &LogComparison{IsSame: false, NeedsReset: true}, nil
	}

	localFile, err := os.Open(localPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open local file: %v", err)
	}
	defer func(localFile *os.File) {
		err = localFile.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(localFile)

	scanner := bufio.NewScanner(localFile)
	compareLineNum := 10
	localLines := make([]string, 0, compareLineNum)
	for i := 0; i < compareLineNum; i++ {
		localLines = append(localLines, scanner.Text())
	}

	remoteReader := bytes.NewReader(remoteContent)
	remoteScanner := bufio.NewScanner(remoteReader)
	remoteLines := make([]string, 0, compareLineNum)
	for i := 0; i < compareLineNum; i++ {
		remoteLines = append(remoteLines, remoteScanner.Text())
	}

	if len(localLines) != len(remoteLines) {
		return &LogComparison{IsSame: false, NeedsReset: true}, nil
	}

	for i := range localLines {
		if localLines[i] != remoteLines[i] {
			return &LogComparison{IsSame: false, NeedsReset: true}, nil
		}
	}
	return &LogComparison{IsSame: true, NeedsReset: false}, nil
}

func fetchLatestLines() {

}
