package git

import "fmt"

func getGitPullRebaseArgs(args ...string) []string {
	fmt.Println("Pulling from remote branch...")
	path := args[0]
	return []string{"git", "-C", path, "pull", "--rebase"}
}

func getGitCheckoutToBranchArgs(args ...string) []string {
	path := args[0]
	branch := args[1]
	fmt.Printf("Switching to branch '%s'...\n", branch)
	return []string{"git", "-C", path, "checkout", branch}
}

func getGitStashArgs(args ...string) []string {
	fmt.Println("Stashing changes...")
	path := args[0]
	return []string{"git", "-C", path, "stash"}
}

func getGitStashPopArgs(args ...string) []string {
	fmt.Println("Poping changes back...")
	path := args[0]
	return []string{"git", "-C", path, "stash", "pop"}
}

func getGitShowCurrentBranchArgs(args ...string) []string {
	path := args[0]
	return []string{"git", "-C", path, "branch", "--show-current"}
}

func getIsGitRepositoryArgs(args ...string) []string {
	path := args[0]
	return []string{"git", "-C", path, "rev-parse", "--is-inside-work-tree"}
}
