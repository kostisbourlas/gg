package git

func getGitPullRebaseArgs(args ...string) []string {
	path := args[0]
	return []string{"git", "-C", path, "pull", "--rebase"}
}

func getGitCheckoutToBranchArgs(args ...string) []string {
	path := args[0]
	branch := args[1]
	return []string{"git", "-C", path, "checkout", branch}
}

func getGitStashArgs(args ...string) []string {
	path := args[0]
	return []string{"git", "-C", path, "stash"}
}

func getGitStashPopArgs(args ...string) []string {
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
