/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>

*/
package git

func getGitPullRebaseArgs(path string)[5]string {
	return [5]string {"git", "-C", path, "pull", "--rebase"}
}

func getGitCheckoutToBranchArgs(path string, branch string)[5]string {
	return [5]string {"git",  "-C", path, "checkout", branch}
}

func getGitStashArgs(path string)[4]string {
	return [4]string {"git", "-C", path, "stash"}
}

func getGitStashPopArgs(path string)[5]string {
	return [5]string {"git", "-C", path, "stash", "pop"}
}

func getGitShowCurrentBranchArgs(path string)[5]string {
	return [5]string {"git", "-C", path, "branch", "--show-current"}
}

func getIsGitRepositoryArgs(path string)[5]string {
	return [5]string {"git", "-C", path, "rev-parse", "--is-inside-work-tree"}
}
