package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	//remove help shorthand for consistency due to clash with high in run
	versionCmd.PersistentFlags().BoolP("help", "", false, "help for this command")

	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ugscli",
	Long:  `Print the version number of ugscli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ugscli - v0.1")
	},
}
