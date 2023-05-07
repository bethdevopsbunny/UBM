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

	getStatusCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	getStatusCmd.PersistentFlags().IntVarP(&buildNumber, "build-number", "b", 0, "build number")
	getStatusCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	getStatusCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")
	getStatusCmd.PersistentFlags().BoolVarP(&raw, "raw-output", "r", false, "returns only raw output")

	getCmd.AddCommand(getStatusCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var getStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		build, _ := req.GetBuildItem(org, buildNumber, buildTarget, projectID, ApiKey)

		response, _ := json.Marshal(build.BuildStatus)
		fmt.Println(string(response))
		os.Exit(0)

	},
}
