package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
	req "ubm/requests"
)

func init() {

	pipelineCmd.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	pipelineCmd.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	pipelineCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")
	pipelineCmd.PersistentFlags().BoolVarP(&raw, "raw-output", "r", false, "returns only raw output")
	pipelineCmd.PersistentFlags().StringVarP(&filePath, "file-path", "f", "artifact.zip", "file path you want the artifact downloaded")
	pipelineCmd.PersistentFlags().DurationVarP(&requestDelay, "request-delay", "d", 20, "delay between requests")

	RootCmd.AddCommand(pipelineCmd)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var success = false
var currentBuildItem req.BuildItem
var err error

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		var buildMean = CalulateMean()

		buildRequest, _ := req.BuildStart(org, buildTarget, projectID, ApiKey)
		sum := 0

	Loop:
		for sum < int(buildMean)+1000 {

			currentBuildItem, err = req.GetBuildItem(org, buildRequest[0].Build, buildTarget, projectID, ApiKey)
			if err != nil {
				log.Warnf("Error: Failed to return build item")
			}

			switch currentBuildItem.BuildStatus {
			case "created":
				log.Infof("Build Created")
			case "queued":
				log.Infof("Building Queued: %s", splitAtUpper(currentBuildItem.QueuedReason))
			case "sentToBuilder":
				log.Infof("Sent to Builder, Waiting for build to start")
			case "started":
				log.Infof("Building...")
			case "restarted":
				log.Infof("Build Restarted")
			case "failure":
				log.Infof("Build Failed")
				break Loop
			case "canceled":
				log.Infof("Build Canceled")
				break Loop
			case "unknown":
				log.Infof("Build Status Unknown, what the hell happened?!")
			case "success":
				log.Infof("Build Complete")
				success = true
				break Loop
			default:
				log.Infof("New Build Status: %s", currentBuildItem.BuildStatus)
				break Loop
			}

			log.Infof("%d", sum)
			time.Sleep(requestDelay * time.Second)
			sum += int(requestDelay)
		}

		if success {
			err = downloadFile(filePath, currentBuildItem.Links.Artifacts[0].Files[0].Href)
			if err != nil {
				log.Errorf("Failed Download... Exiting")
				os.Exit(1)
			}
		}

		os.Exit(0)
	},
}
