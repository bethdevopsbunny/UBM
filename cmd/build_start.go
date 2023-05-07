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

	createCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	createCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	createCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")
	createCmd.PersistentFlags().BoolVarP(&raw, "raw-output", "r", false, "returns only raw output")

	buildCmd.AddCommand(createCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("create build")

		authentication()

		build, _ := req.BuildStart(org, buildTarget, projectID, ApiKey)

		response, _ := json.Marshal(build)
		fmt.Println(string(response))
		os.Exit(0)

	},
}
