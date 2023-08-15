package cmd

import (
	"fmt"
	Git "github.com/kostisbourlas/gg/git"
	"github.com/spf13/cobra"
)

/*
Usage:
1. gg preview-merge master feature_branch --path /home/$USER/directory/repo
2. gg preview-merge master feature_branch
*/

var previewMergeCmd = &cobra.Command{
	Use:   "preview-merge",
	Short: "Previews the result of a merge.",
	Long:  `Previews the result of a merge.`,
	Run:   previewMergeRun,
}

func previewMergeRun(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Two branches should be given.")
		return
	}

	var branches []string
	for _, arg := range args {
		branches = append(branches, arg)
	}

	path, _ := cmd.Flags().GetString("path")
	if path == "" {
		path = getCurrentDir()
	}

	// checks out to branch that the merge will be performed.
	err := Git.CheckoutToBranch(path, branches[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Git.PreviewMerge(path, branches[1])
	if err != nil {
		fmt.Println(err)
		return
	}

}

func init() {
	rootCmd.AddCommand(previewMergeCmd)
	previewMergeCmd.Flags().StringP("path", "p", "", "path of the repo")
}
