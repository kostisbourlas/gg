package git

import (
	"fmt"
	"os/exec"
)

func performGitStashPop(path string) ([]byte, error) {
	fmt.Println("Performing git stash pop...")
	gitStashPopArgs := getGitStashPopArgs(path)
	gitStashPopCmd := exec.Command(gitStashPopArgs[0], gitStashPopArgs[1:]...)
	output, err := gitStashPopCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}

func performGitStash(path string) ([]byte, error) {
	fmt.Println("Performing git stash...")
	gitStashArgs := getGitStashArgs(path)
	gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
	output, err := gitStashCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}

func performGitPull(path string) ([]byte, error) {
	fmt.Println("Performing git pull --rebase...")
	gitPullRebaseArgs := getGitPullRebaseArgs(path)
	gitPullRebaseCmd := exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
	output, err := gitPullRebaseCmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}

func performGitIsRepository(path string) ([]byte, error) {
	cmdArgs := getIsGitRepositoryArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	return output, err
}

func performGitCurrentBranch(path string) ([]byte, error) {
	cmdArgs := getGitShowCurrentBranchArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	return output, err
}

func performGitCheckoutToBranch(path string, branch string) ([]byte, error) {
	fmt.Printf("Performing git checkout to %s...\n", branch)
	cmdArgs := getGitCheckoutToBranchArgs(path, branch)
	gitCheckoutCommand := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := gitCheckoutCommand.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}
