package cmd

import (
	"github.com/creack/pty"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
)

//*/
//	Usage:
//	1. gg update /home/$USER/directory/repo1 --branch devel
//	2. gg update ~/directory/repo1 ~/directory/repo2 --branch devel
//	3. gg update ~/directory/repo1
//*/

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows a git repo",
	Long:  `This command shows a  given git repo.`,
	Run:   showRun,
}

func showRun(cmd *cobra.Command, args []string) {
	StatusCmd := exec.Command("git", "-C", "/home/kostis/Projects/txc", "status")

	f, err := pty.Start(StatusCmd)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)
}

func init() {
	rootCmd.AddCommand(showCmd)
}
