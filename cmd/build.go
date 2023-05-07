package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {

	RootCmd.AddCommand(buildCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("build options")
	},
}
