package git

import (
	"fmt"
	"strings"
)

func IsGitRepository(path string) bool {
	output, err := executeGitCommand(getIsGitRepositoryArgs, path)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func GetCurrentBranch(path string) string {
	output, _ := executeGitCommand(getGitShowCurrentBranchArgs, path)
	return strings.TrimSpace(string(output))
}

func CheckoutToBranch(path string, branch string) error {
	_, err := executeGitCommand(getGitCheckoutToBranchArgs, path, branch)

	// performs git stash first in order to check out successfully
	if err != nil {
		_, err = executeGitCommand(getGitStashArgs, path)
		if err != nil {
			return fmt.Errorf("git stash failed with error: %v", err)
		}

		_, err = executeGitCommand(getGitCheckoutToBranchArgs, path, branch)
		if err != nil {
			return fmt.Errorf("git checkout to %s failed with error: %v", branch, err)
		}
	}
	return nil
}

func UpdateGitRepository(path string) error {
	_, err := executeGitCommand(getGitPullRebaseArgs, path)

	// performs git stash first in order to git pull successfully
	if err != nil {
		_, err = executeGitCommand(getGitStashArgs, path)
		if err != nil {
			return fmt.Errorf("git stash failed with error: %v", err)
		}

		_, err = executeGitCommand(getGitPullRebaseArgs, path)
		if err != nil {
			return fmt.Errorf("git pull --rebase failed with error: %v", err)
		}

		_, err = executeGitCommand(getGitStashPopArgs, path)
		if err != nil {
			return fmt.Errorf("git stash pop failed with error: %v", err)
		}
	}
	return nil
}
