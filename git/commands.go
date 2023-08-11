package git

import (
	"fmt"
	"os/exec"
)

func performGitStashPop(path string) {
	fmt.Println("Performing git stash pop...")
	gitStashPopArgs := getGitStashPopArgs(path)
	gitStashPopCmd := exec.Command(gitStashPopArgs[0], gitStashPopArgs[1:]...)
	output, _ := gitStashPopCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
}

func performGitStash(path string) {
	fmt.Println("Performing git stash...")
	gitStashArgs := getGitStashArgs(path)
	gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
	output, _ := gitStashCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
}

func performGitPull(path string) error {
	fmt.Println("Performing git pull --rebase...")
	gitPullRebaseArgs := getGitPullRebaseArgs(path)
	gitPullRebaseCmd := exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
	output, err := gitPullRebaseCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return err
}
