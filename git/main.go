/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>
*/
package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsGitRepository(path string) bool {
	cmdArgs := getIsGitRepositoryArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func GetCurrentBranch(path string) string {
	cmdArgs := getGitShowCurrentBranchArgs(path)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, _ := cmd.CombinedOutput()
	return strings.TrimSpace(string(output))
}

func CheckoutToBranch(path string, branch string) error {
	fmt.Printf("Performing git checkout to %s...\n", branch)
	cmdArgs := getGitCheckoutToBranchArgs(path, branch)
	gitCheckoutCommand := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := gitCheckoutCommand.CombinedOutput()
	fmt.Printf("%s\n", output)

	// performs git stash first if cannot check out to branch
	if err != nil {
		fmt.Println("Performing git stash...")
		gitStashArgs := getGitStashArgs(path)
		gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
		output, _ := gitStashCmd.CombinedOutput()
		fmt.Printf("%s\n", output)

		gitCheckoutCommand = exec.Command(cmdArgs[0], cmdArgs[1:]...)
		output, err = gitCheckoutCommand.CombinedOutput()
		fmt.Printf("%s\n", output)

		if err != nil {
			return fmt.Errorf("git checkout to %s failed with error: %v", branch, err)
		}
	}
	return nil
}

func UpdateGitRepository(path string) error {
	fmt.Println("Performing git pull --rebase...")
	gitPullRebaseArgs := getGitPullRebaseArgs(path)
	gitPullRebaseCmd := exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
	output, err := gitPullRebaseCmd.CombinedOutput()
	fmt.Printf("%s\n", output)

	// performs git stash first in order to git pull successfully
	if err != nil {
		fmt.Println("Performing git stash...")
		gitStashArgs := getGitStashArgs(path)
		gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
		output, _ = gitStashCmd.CombinedOutput()
		fmt.Printf("%s\n", output)

		fmt.Println("Performing git pull --rebase...")
		gitPullRebaseCmd = exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
		output, err = gitPullRebaseCmd.CombinedOutput()
		fmt.Printf("%s\n", output)
		if err != nil {
			return fmt.Errorf("git pull --rebase failed with error: %v", err)
		}

		fmt.Println("Performing git stash pop...")
		gitStashPopArgs := getGitStashPopArgs(path)
		gitStashPopCmd := exec.Command(gitStashPopArgs[0], gitStashPopArgs[1:]...)
		output, _ = gitStashPopCmd.CombinedOutput()
		fmt.Printf("%s\n", output)
	}
	return nil
}
