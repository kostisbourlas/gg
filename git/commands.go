package git

import (
	"fmt"
	"os/exec"
)

type gitParams func(...string) []string

func executeGitCommand(getGitCommandParams gitParams, args ...string) ([]byte, error) {
	var cmdArgs []string
	if len(args) == 1 {
		cmdArgs = getGitCommandParams(args[0])
	} else if len(args) == 2 {
		cmdArgs = getGitCommandParams(args[0], args[1])
	} else {
		fmt.Println("Invalid number of parameters. Give one or two.")
	}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}
