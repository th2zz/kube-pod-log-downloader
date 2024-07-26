package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getPods(regex string) ([]string, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-o", "jsonpath={.items[*].metadata.name}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error getting pods: %v", err)
	}

	pods := strings.Fields(out.String())
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("invalid regex: %v", err)
	}

	var filteredPods []string
	for _, pod := range pods {
		if pattern.MatchString(pod) {
			filteredPods = append(filteredPods, pod)
		}
	}
	return filteredPods, nil
}

func captureLogs(pods []string) {
	for _, pod := range pods {
		logFile := fmt.Sprintf("%s.log", pod)
		cmd := exec.Command("kubectl", "logs", pod)
		out, err := os.Create(logFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating log file for pod %s: %v\n", pod, err)
			continue
		}
		defer out.Close()
		cmd.Stdout = out
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error capturing logs for pod %s: %v\n", pod, err)
		} else {
			fmt.Printf("Captured logs for pod %s in %s\n", pod, logFile)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run script.go <pod_name_regex>")
		os.Exit(1)
	}

	regex := os.Args[1]
	pods, err := getPods(regex)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(pods) == 0 {
		fmt.Fprintf(os.Stderr, "No pods found matching regex %s\n", regex)
		os.Exit(1)
	}

	captureLogs(pods)
}
