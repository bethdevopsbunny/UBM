package requests

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type BuildAccepted struct {
	Build           int           `json:"build"`
	Buildtargetid   string        `json:"buildtargetid"`
	BuildTargetName string        `json:"buildTargetName"`
	BuildStatus     string        `json:"buildStatus"`
	CleanBuild      bool          `json:"cleanBuild"`
	Platform        string        `json:"platform"`
	Created         time.Time     `json:"created"`
	Changeset       []interface{} `json:"changeset"`
	Favorited       bool          `json:"favorited"`
	AuditChanges    int           `json:"auditChanges"`
	Links           struct {
		Self struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"self"`
		Log struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"log"`
		Auditlog struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"auditlog"`
		Cancel struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"cancel"`
	} `json:"links"`
}

type BuildRequestFailed struct {
	Buildtargetid   string `json:"buildtargetid"`
	BuildTargetName string `json:"buildTargetName"`
	Error           string `json:"error"`
}

func BuildStart(org int, buildTarget string, projectID string, key string) ([]BuildAccepted, error) {

	var buildAccepted []BuildAccepted
	var buildRequestFailed []BuildRequestFailed
	path := fmt.Sprintf("https://build-api.cloud.unity3d.com/api/v1/orgs/%d/projects/%s/buildtargets/%s/builds", org, projectID, buildTarget)
	auth := fmt.Sprintf("Basic: %s", key)

	data := url.Values{}
	//data.Set("clean", "true")
	//data.Set("delay", "30")

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &buildAccepted)

	// My understanding is that when you already have a build in queued and make another request your build might still be accepted.
	// however you are returned a buildRequestFailed object instead with an error and no build number.
	// im not 100% but i think it caches this request and then runs it in its own time. This SUCKS.
	// breaks this workflow and unless i can think of a way to overcome this might build in a failsafe to prevent a request when thier is a build in that state
	// in this app.

	// because failures still return status code 202 we have to build the accepted object and check if its nil
	// the build number seemed like a good choice here.
	if buildAccepted[0].Build == 0 {

		err = json.Unmarshal(body, &buildRequestFailed)
		log.Errorf("Build Request Failed: %s", buildRequestFailed[0].Error)

		os.Exit(1)

	}

	return buildAccepted, nil

}
