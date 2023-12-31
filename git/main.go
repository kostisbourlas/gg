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

func PreviewMerge(path string, branch string) error {
	/*
		Performs the following commands in order to preview the conflicts if any.
			git merge <branch> --no-ff --no-commit
			git diff <branch>
			git merge --abort
		If there are no conflicts instead of git diff andgit  merge --abort, performs:
			git reset --hard HEAD
	*/
	output, _ := executeGitCommand(getGitMergeArgs, path, branch)
	if !strings.Contains(string(output), "Automatic merge went well") {
		_, _ = executeGitCommand(getGitResetHard, path)
		return nil
	}

	_, err := executeGitCommand(getGitDiffArgs, path, branch)
	if err != nil {
		return fmt.Errorf("git diff failed with error: %v", err)
	}

	_, err = executeGitCommand(getGitMergeAbortArgs, path, branch)

	return nil
}
