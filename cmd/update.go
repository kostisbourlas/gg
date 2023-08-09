/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	Git "github.com/kostisbourlas/gg/git"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a git repo",
	Long: `This command updates a  given git repo.`,
	Run: updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	var paths []string
	for _, arg := range args {
		paths = append(paths, arg)
	}
	branch, _ := cmd.Flags().GetString("branch")

	for _, path := range paths {
		isGitRepo := Git.IsGitRepository(path)
		if isGitRepo == false {
			fmt.Println("Path is not a git repository.")
			return 
		}
	
		current_branch := Git.GetCurrentBranch(path)
	
		err := Git.CheckoutToBranch(path, branch)
		if err != nil {
			fmt.Println(err)
			return
		}
	
		err = Git.UpdateGitRepository(path)
		if err != nil {
			fmt.Println(err)
			return 
		}
	
		err = Git.CheckoutToBranch(path, current_branch)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Update ran successfully!")
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("branch", "b", "master", "specific git branch")
}
