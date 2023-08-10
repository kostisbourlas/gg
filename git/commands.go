
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
	branch := strings.TrimSpace(string(output))
	return branch
}

func CheckoutToBranch(path string, branch string) error {
	cmdArgs := getGitCheckoutToBranchArgs(path, branch)
	gitCheckoutCommand := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	output, err := gitCheckoutCommand.CombinedOutput()

	// performs git stash first if cannot checkout to branch
	if err != nil {
		gitStashArgs := getGitStashArgs(path)
		gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
		_, err := gitStashCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("checking out to branch: %s failed with error: %v", branch, err)
		}
		gitCheckoutCommand = exec.Command(cmdArgs[0], cmdArgs[1:]...)
		output, err := gitCheckoutCommand.CombinedOutput()
		fmt.Printf("%s\n", output)
		return nil
	}
	fmt.Printf("%s\n", output)
	return nil
}

func UpdateGitRepository(path string) error {
	gitPullRebaseArgs := getGitPullRebaseArgs(path)

	gitPullRebaseCmd := exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
	output, err := gitPullRebaseCmd.CombinedOutput()

	// performs git stash first in order to git pull successfully
	if err != nil {
		fmt.Println("Performing git stash...")
		gitStashArgs := getGitStashArgs(path)
		gitStashCmd := exec.Command(gitStashArgs[0], gitStashArgs[1:]...)
		_, err := gitStashCmd.CombinedOutput()
		if err != nil {
			return 	fmt.Errorf("error upgrading Git repository: %v", err)
		}
		gitPullRebaseCmd := exec.Command(gitPullRebaseArgs[0], gitPullRebaseArgs[1:]...)
		output, _ := gitPullRebaseCmd.CombinedOutput()
		fmt.Printf("%s\n", output)

		gitStashPopArgs := getGitStashPopArgs(path)
		gitStashPopCmd := exec.Command(gitStashPopArgs[0], gitStashPopArgs[1:]...)
		output, _ = gitStashPopCmd.CombinedOutput()
		fmt.Printf("%s\n", output)
		return nil
	}
	fmt.Printf("%s\n", output)
	return nil
}
