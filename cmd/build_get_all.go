package cmd

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	req "ubm/requests"
)

func init() {

	getAllCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	getAllCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	getAllCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")

	getCmd.AddCommand(getAllCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var getAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		build, _ := req.GetBuildList(org, buildTarget, projectID, ApiKey, 100, true)

		response, _ := json.Marshal(build)
		fmt.Println(string(response))
		os.Exit(0)

	},
}
