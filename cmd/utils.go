package cmd

import (
	"os/exec"
	"strings"
)

func getCurrentDir() string {
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output)) + "/"
}
