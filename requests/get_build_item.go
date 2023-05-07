package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BuildItem struct {
	Build                 int           `json:"build"`
	Buildtargetid         string        `json:"buildtargetid"`
	BuildTargetName       string        `json:"buildTargetName"`
	BuildGUID             string        `json:"buildGUID"`
	BuildStatus           string        `json:"buildStatus"`
	CleanBuild            bool          `json:"cleanBuild"`
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
	BillableTimeInSeconds int           `json:"billableTimeInSeconds"`
	UnitTestTimeInSeconds int           `json:"unitTestTimeInSeconds"`
	LastBuiltRevision     string        `json:"lastBuiltRevision"`
	Changeset             []interface{} `json:"changeset"`
	Favorited             bool          `json:"favorited"`
	Deleted               bool          `json:"deleted"`
	Headless              bool          `json:"headless"`
	CredentialsOutdated   bool          `json:"credentialsOutdated"`
	QueuedReason          string        `json:"queuedReason"`
	CooldownDate          time.Time     `json:"cooldownDate"`
	ScmBranch             string        `json:"scmBranch"`
	UnityVersion          string        `json:"unityVersion"`
	LocalUnityVersion     string        `json:"localUnityVersion"`
	AuditChanges          int           `json:"auditChanges"`
	ProjectVersion        struct {
		Name        string        `json:"name"`
		Filename    string        `json:"filename"`
		ProjectName string        `json:"projectName"`
		Platform    string        `json:"platform"`
		Size        int           `json:"size"`
		Created     time.Time     `json:"created"`
		LastMod     time.Time     `json:"lastMod"`
		Udids       []interface{} `json:"udids"`
	} `json:"projectVersion"`
	ProjectName string `json:"projectName"`
	ProjectId   string `json:"projectId"`
	ProjectGuid string `json:"projectGuid"`
	OrgId       string `json:"orgId"`
	OrgFk       string `json:"orgFk"`
	Filetoken   string `json:"filetoken"`
	Links       struct {
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
		DownloadPrimary struct {
			Method string `json:"method"`
			Href   string `json:"href"`
			Meta   struct {
				Type string `json:"type"`
			} `json:"meta"`
		} `json:"download_primary"`
		Artifacts []struct {
			Key           string `json:"key"`
			Name          string `json:"name"`
			Primary       bool   `json:"primary"`
			ShowDownload  bool   `json:"show_download"`
			WebglTemplate string `json:"webgl_template"`
			Files         []struct {
				Filename  string `json:"filename"`
				Size      int    `json:"size"`
				Resumable bool   `json:"resumable"`
				Md5Sum    string `json:"md5sum"`
				Href      string `json:"href"`
			} `json:"files"`
		} `json:"artifacts"`
	} `json:"links"`
	BuildReport struct {
		Errors   int `json:"errors"`
		Warnings int `json:"warnings"`
	} `json:"buildReport"`
}

func GetBuildItem(org int, buildNumber int, buildTarget string, projectID string, key string) (BuildItem, error) {

	var buildItem BuildItem
	url := fmt.Sprintf("https://build-api.cloud.unity3d.com/api/v1/orgs/%d/projects/%s/buildtargets/%s/builds/%d", org, projectID, buildTarget, buildNumber)
	auth := fmt.Sprintf("Basic: %s", key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BuildItem{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return BuildItem{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return BuildItem{}, err
	}

	err = json.Unmarshal(body, &buildItem)
	if err != nil {
		return BuildItem{}, err
	}

	return buildItem, nil
}
