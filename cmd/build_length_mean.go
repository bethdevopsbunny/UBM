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

	buildLengthMean.PersistentFlags().IntVarP(&org, "org", "o", 0, "organisation number")
	buildLengthMean.PersistentFlags().StringVarP(&buildTarget, "build-target", "t", "", "build target")
	buildLengthMean.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "project id")

	buildCmd.AddCommand(buildLengthMean)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

}

var totalTimesInSeconds []float64

var buildLengthMean = &cobra.Command{
	Use:   "mean",
	Short: "Unity Build Manager",
	Long:  " \nCLI for managing Unity Gaming Services Builds",
	Run: func(cmd *cobra.Command, args []string) {

		authentication()

		response, _ := json.Marshal(CalulateMean())
		fmt.Println(string(response))
		os.Exit(0)

	},
}

func CalulateMean() float64 {
	buildItems, _ := req.GetBuildList(org, buildTarget, projectID, ApiKey, 100, true)

	for _, item := range buildItems {
		totalTimesInSeconds = append(totalTimesInSeconds, item.TotalTimeInSeconds)
	}
	return mean(totalTimesInSeconds)
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}
