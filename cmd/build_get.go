spackage cmd

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	req "ubm/requests"
)

func init() {

	getCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	getCmd.PersistentFlags().IntVarP(&buildNumber, "build-number", "b", 0, "build number")
	getCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	getCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")
	getCmd.PersistentFlags().BoolVarP(&raw, "raw-output", "r", false, "returns only raw output")

	buildCmd.AddCommand(getCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		build, _ := req.GetBuildItem(org, buildNumber, buildTarget, projectID, ApiKey)

		if raw {
			response, _ := json.Marshal(build)
			fmt.Println(string(response))
			os.Exit(0)
		}

		response, _ := json.Marshal(
			returnValue{
				build.BuildStatus,
				build.BuildGUID,
			},
		)
		fmt.Println(string(response))
		os.Exit(0)

	},
}
