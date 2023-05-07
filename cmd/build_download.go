package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	req "ubm/requests"
)

func init() {

	downloadCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	downloadCmd.PersistentFlags().IntVarP(&buildNumber, "build-number", "b", 0, "build number")
	downloadCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	downloadCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")
	downloadCmd.PersistentFlags().StringVarP(&filePath, "file-path", "f", "artifact.zip", "file path you want the artifact downloaded")

	buildCmd.AddCommand(downloadCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		buildItem, _ := req.GetBuildItem(org, buildNumber, buildTarget, projectID, ApiKey)
		downloadFile(filePath, buildItem.Links.Artifacts[0].Files[0].Href)

		os.Exit(0)

	},
}
