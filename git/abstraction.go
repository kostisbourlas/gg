package git

import (
	"fmt"
	"os/exec"
)

type gitParams func(...string) []string

func executeGitCommand(getGitCommandParams gitParams, args ...string) ([]byte, error) {
	cmdArgs := extractGitArgs(getGitCommandParams, args)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}
