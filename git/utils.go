
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
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}

func GetCurrentBranch(path string) string {
	cmd := exec.Command("git", "-C", path, "branch", "--show-current")
	output, _ := cmd.CombinedOutput()
	branch := strings.TrimSpace(string(output))
	return branch
}

func CheckoutToBranch(path string, branch string) error {
	cmd := exec.Command("git",  "-C", path, "checkout", branch)
	output, err := cmd.CombinedOutput()

	// performs git stash first if cannot checkout to branch
	if err != nil {
		cmd := exec.Command("git", "-C", path, "stash")
		_, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("checking out to branch: %s failed with error: %v", branch, err)
		}
		cmd = exec.Command("git",  "-C", path, "checkout", branch)
		output, err := cmd.CombinedOutput()
		fmt.Printf("%s\n", output)
		return nil
	}
	fmt.Printf("%s\n", output)
	return nil
}

func UpdateGitRepository(path string) error {
	cmd := exec.Command("git", "-C", path, "pull", "--rebase")
	output, err := cmd.CombinedOutput()

	// performs git stash first in order to git pull successfully
	if err != nil {
		fmt.Println("Performing git stash...")
		cmd := exec.Command("git", "-C", path, "stash")
		_, err := cmd.CombinedOutput()
		if err != nil {
			return 	fmt.Errorf("error upgrading Git repository: %v", err)
		}
		cmd = exec.Command("git", "-C", path, "pull", "--rebase")
		output, _ := cmd.CombinedOutput()
		fmt.Printf("%s\n", output)

		cmd = exec.Command("git", "-C", path, "stash", "pop")
		return nil
	}
	fmt.Printf("%s\n", output)
	return nil
}
