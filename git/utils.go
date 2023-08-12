package git

import (
	"fmt"
	"github.com/creack/pty"
	"os/exec"
)

func extractGitArgs(getGitCommandParams gitParams, args []string) []string {
	var cmdArgs []string
	if len(args) == 1 {
		cmdArgs = getGitCommandParams(args[0])
	} else if len(args) == 2 {
		cmdArgs = getGitCommandParams(args[0], args[1])
	} else {
		fmt.Println("Invalid number of parameters. Give one or two.")
	}
	return cmdArgs
}

func executeCmdInPseudoTerminal(cmd *exec.Cmd) ([]byte, error) {
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}
	defer ptmx.Close()

	outputBytes := make([]byte, 0)
	for {
		var n int
		buf := make([]byte, 4096)
		n, err = ptmx.Read(buf)
		if err != nil {
			break
		}
		outputBytes = append(outputBytes, buf[:n]...)
	}

	err = cmd.Wait()
	return outputBytes, err
}
