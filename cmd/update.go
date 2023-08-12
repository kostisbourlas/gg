package cmd

import (
	"fmt"

	Git "github.com/kostisbourlas/gg/git"
	"github.com/spf13/cobra"
)

/*
Usage:
1. gg update /home/$USER/directory/repo1 --branch devel
2. gg update ~/directory/repo1 ~/directory/repo2 --branch devel
3. gg update ~/directory/repo1
*/

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a git repo",
	Long:  `This command updates a  given git repo.`,
	Run:   updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	var paths []string
	for _, arg := range args {
		paths = append(paths, arg)
	}
	branch, _ := cmd.Flags().GetString("branch")

	for _, path := range paths {
		fmt.Printf("Initializing update for: %s \n", path)

		isGitRepo := Git.IsGitRepository(path)
		if isGitRepo == false {
			fmt.Println("Path is not a git repository.")
			continue
		}

		currentBranch := Git.GetCurrentBranch(path)

		err := Git.CheckoutToBranch(path, branch)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = Git.UpdateGitRepository(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = Git.CheckoutToBranch(path, currentBranch)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("branch", "b", "master", "specific git branch")
}
