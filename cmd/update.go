/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>

*/
package cmd

import (
	"fmt"
	"os"
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
	
	err := changeDirectory(repo)
	if err != nil {
		fmt.Println(err)
	}
	isGitRepo := isGitRepository(repo)
	
	if isGitRepo == false {
		fmt.Println("Given path is not a git repository.")
		return 
	}

	err = checkoutToBranch(repo, branch)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = updateGitRepository(repo)
	if err != nil {
		fmt.Println(err)
		return 

	}
}

func changeDirectory(path string) error {
	// checks if directory exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", path)
	}

	// Change the working directory
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("error changing directory: %v", err)
	}
	return nil
}

func isGitRepository(path string) bool {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}

func checkoutToBranch(path string, branch string) error {
	cmd := exec.Command("git",  "-C", path, "checkout", branch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error checking out to branch %s: %v", branch, err)
	}
	fmt.Printf("%s\n", output)
	return nil
}

func updateGitRepository(path string) error {
	cmd := exec.Command("git", "-C", path, "pull", "--rebase")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 	fmt.Errorf("error upgrading Git repository: %v", err)
	}
	fmt.Printf("%s\n", output)
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("repo", "r", "", "specific git repository")
	updateCmd.Flags().StringP("branch", "b", "", "specific git branch")
}
