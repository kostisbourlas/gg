/*
Copyright © 2023 Kostis Bourlas <kostisbourlas@protonmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

/*
	gg update --path /home/kostis/Projects/ --branch devel
	gg update --repo txc/
*/

var rootCmd = &cobra.Command{
	Use:   "gg",
	Short: "A Git Repository Manager",
	Long: "A Git Repository Manager",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}


