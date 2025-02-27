package fetch

import (
	"LogSentinel/config"
	"LogSentinel/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func K8sPodLogs() {
	logger := utils.GetLogger()

	podLogs := getPodLogs()
	if len(podLogs) == 0 {
		logger.Warn("No pods found for log retrieval.")
		return
	}

	var wg sync.WaitGroup
	for _, podLog := range podLogs {
		wg.Add(1)
		go func(podLog *PodLog) {
			defer wg.Done()
			logger.Info(fmt.Sprintf("Fetching logs for pod: %s (namespace: %s)", podLog.PodName, podLog.Namespace))

			var cmd *exec.Cmd
			if podLog.Container != "" {
				cmd = exec.Command("kubectl", "exec", "-i", podLog.PodName, "-n", podLog.Namespace, "-c", podLog.Container, "--", "sed", "-n", "1,100p", podLog.LogPath)
			} else {
				cmd = exec.Command("kubectl", "exec", "-i", podLog.PodName, "-n", podLog.Namespace, "--", "sed", "-n", "1,100p", podLog.LogPath)
			}

			output, err := cmd.CombinedOutput()

			if err != nil {
				logger.Error(fmt.Sprintf("Failed to fetch logs for pod %s: %s", podLog.PodName, err.Error()))
				return
			}

			logContent := strings.TrimSpace(string(output))
			if logContent == "" {
				logger.Warn(fmt.Sprintf("No logs found for pod %s", fmt.Sprintf("%s:%s", podLog.PodName, podLog.LogPath)))
				return
			}

			logger.Debug(fmt.Sprintf("Logs for pod %s:\n%s", podLog.PodName, logContent))
			writeFetchedLogToDisk(podLog.FetchDest, logContent)
		}(podLog)
	}
	wg.Wait()
}

func getPodLogs() []*PodLog {
	logger := utils.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil
	}

	var pods []*PodLog

	for _, kube := range cfg.Kube {
		logger.Info("Fetching Pods from %s", kube.Name)
		getPodsCmd := exec.Command("kubectl", "get", "pods", "-n", kube.NameSpace, "-l", kube.ServiceLabel, "--no-headers")
		output, err := getPodsCmd.Output()
		if err != nil {
			logger.Error(err.Error())
			return nil
		}
		podList := strings.Split(strings.TrimSpace(string(output)), "\n")
		if len(podList) > 1 {
			logger.Info("Found %d Pods", len(podList))
			logger.Info("Selecting specific pod from list isn't implemented yet. It will automatically select the first Pod")
		}
		podName := strings.Fields(podList[0])[0]
		fetchDest := kube.Name + "." + kube.Format
		podLog := PodLog{
			PodName:   podName,
			Namespace: kube.NameSpace,
			LogPath:   kube.LogPath,
			FetchDest: fetchDest,
			Container: kube.Container,
		}
		pods = append(pods, &podLog)
	}
	return pods

}

func writeFetchedLogToDisk(dest string, content string) {
	logger := utils.GetLogger()

	path := "./remote_logs/" + dest
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer func(logFile *os.File) {
		err = logFile.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(logFile)

	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

	_, err = logFile.WriteString(content)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
