/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>

*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a git repo",
	Long: `This command updates a  given git repo.`,
	Run: updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	repo, _ := cmd.Flags().GetString("repo")
	branch, _ := cmd.Flags().GetString("branch")
	
	isGitRepo := isGitRepository(repo)
	if isGitRepo == false {
		fmt.Println("Path is not a git repository.")
		return 
	}

	current_branch := getCurrentBranch(repo)

	err := checkoutToBranch(repo, branch)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = updateGitRepository(repo)
	if err != nil {
		fmt.Println(err)
		return 
	}

	err = checkoutToBranch(repo, current_branch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update ran successfully!")
}

func isGitRepository(path string) bool {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}

func getCurrentBranch(path string) string {
	cmd := exec.Command("git", "-C", path, "branch", "--show-current")
	output, _ := cmd.CombinedOutput()
	branch := strings.TrimSpace(string(output))
	return branch
}

func checkoutToBranch(path string, branch string) error {
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

func updateGitRepository(path string) error {
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

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("repo", "r", "", "specific git repository")
	updateCmd.Flags().StringP("branch", "b", "", "specific git branch")
}
