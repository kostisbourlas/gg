package git

import (
	"fmt"
	"os/exec"
)

func performGitStashPop(path string) ([]byte, error) {
	fmt.Println("Performing git stash pop...")
	cmdArgs := getGitStashPopArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}

func performGitStash(path string) ([]byte, error) {
	fmt.Println("Performing git stash...")
	cmdArgs := getGitStashArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}

func performGitPull(path string) ([]byte, error) {
	fmt.Println("Performing git pull --rebase...")
	cmdArgs := getGitPullRebaseArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
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
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)
	return output, err
}
