package git

import (
	"fmt"
	"os/exec"
)

type gitParams func(...string) []string

func executeGitCommand(getGitCommandParams gitParams, args ...string) ([]byte, error) {
	cmdArgs := extractGitArgs(getGitCommandParams, args)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	outputBytes, err := executeCmdInPseudoTerminal(cmd)
	fmt.Printf("%s", outputBytes)
	return outputBytes, err
}
