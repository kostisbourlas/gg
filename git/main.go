/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>
*/

package git

import (
	"fmt"
	"strings"
)

func IsGitRepository(path string) bool {
	output, err := performGitIsRepository(path)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

func GetCurrentBranch(path string) string {
	output, _ := performGitCurrentBranch(path)
	return strings.TrimSpace(string(output))
}

func CheckoutToBranch(path string, branch string) error {
	_, err := performGitCheckoutToBranch(path, branch)

	// performs git stash first if it cannot check out to branch
	if err != nil {
		_, err = performGitStash(path)
		if err != nil {
			return fmt.Errorf("git stash failed with error: %v", err)
		}

		_, err = performGitCheckoutToBranch(path, branch)
		if err != nil {
			return fmt.Errorf("git checkout to %s failed with error: %v", branch, err)
		}
	}
	return nil
}

func UpdateGitRepository(path string) error {
	_, err := performGitPull(path)

	// performs git stash first in order to git pull successfully
	if err != nil {
		_, err = performGitStash(path)
		if err != nil {
			return fmt.Errorf("git stash failed with error: %v", err)
		}

		_, err = performGitPull(path)
		if err != nil {
			return fmt.Errorf("git pull --rebase failed with error: %v", err)
		}

		_, err = performGitStashPop(path)
		if err != nil {
			return fmt.Errorf("git stash pop failed with error: %v", err)
		}
	}
	return nil
}
