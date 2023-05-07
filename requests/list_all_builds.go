package requests

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type BuildListItem struct {
	OrgId                 interface{}   `json:"orgId"`
	ProjectId             interface{}   `json:"projectId"`
	ProjectName           interface{}   `json:"projectName"`
	Build                 int           `json:"build"`
	Buildtargetid         string        `json:"buildtargetid"`
	BuildTargetName       string        `json:"buildTargetName"`
	BuildStatus           string        `json:"buildStatus"`
	Platform              string        `json:"platform"`
	WorkspaceSize         int           `json:"workspaceSize"`
	Created               time.Time     `json:"created"`
	Finished              time.Time     `json:"finished"`
	CheckoutStartTime     time.Time     `json:"checkoutStartTime"`
	CheckoutTimeInSeconds int           `json:"checkoutTimeInSeconds"`
	BuildStartTime        time.Time     `json:"buildStartTime"`
	BuildTimeInSeconds    float64       `json:"buildTimeInSeconds"`
	PublishStartTime      time.Time     `json:"publishStartTime"`
	PublishTimeInSeconds  float64       `json:"publishTimeInSeconds"`
	TotalTimeInSeconds    float64       `json:"totalTimeInSeconds"`
	LastBuiltRevision     string        `json:"lastBuiltRevision"`
	Changeset             []interface{} `json:"changeset"`
	Favorited             bool          `json:"favorited"`
	ScmBranch             string        `json:"scmBranch"`
	UnityVersion          string        `json:"unityVersion"`
	AuditChanges          int           `json:"auditChanges"`
	ProjectVersion        struct {
		Name        string        `json:"name"`
		Filename    string        `json:"filename"`
		ProjectName string        `json:"projectName"`
		Platform    string        `json:"platform"`
		Size        int           `json:"size"`
		Created     time.Time     `json:"created"`
		LastMod     time.Time     `json:"lastMod"`
		BundleId    string        `json:"bundleId"`
		Udids       []interface{} `json:"udids"`
	} `json:"projectVersion"`
	Links struct {
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
		DownloadPrimary struct {
			Method string `json:"method"`
			Href   string `json:"href"`
			Meta   struct {
				Type string `json:"type"`
			} `json:"meta"`
		} `json:"download_primary"`
		CreateShare struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"create_share"`
		RevokeShare struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"revoke_share"`
		Icon struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"icon"`
	} `json:"links"`
	//CleanBuild string `json:"cleanBuild,omitempty"`
}

func GetBuildList(org int, buildTarget string, projectID string, key string, count int, onlySuccessful bool) ([]BuildListItem, error) {

	var buildList []BuildListItem
	url := fmt.Sprintf("https://build-api.cloud.unity3d.com/api/v1/orgs/%d/projects/%s/buildtargets/%s/builds?per_page=%d", org, projectID, buildTarget, count)
	if onlySuccessful {
		url = fmt.Sprintf("%s&buildStatus=success", url)
	}
	auth := fmt.Sprintf("Basic: %s", key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []BuildListItem{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []BuildListItem{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []BuildListItem{}, err
	}

	err = json.Unmarshal(body, &buildList)
	if err != nil {
		log.Warnf("errorr")
	}

	return buildList, nil
}
