package git

import "fmt"

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
