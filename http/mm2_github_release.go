package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"net/http"
	"runtime"
	"strings"
	"time"
)

const targetUrl = "https://api.github.com/repos/KomodoPlatform/komodo-defi-framework/releases/latest"

type GithubLatestRelease struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string      `json:"node_id"`
	TagName         string      `json:"tag_name"`
	TargetCommitish string      `json:"target_commitish"`
	Name            interface{} `json:"name"`
	Draft           bool        `json:"draft"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       time.Time   `json:"created_at"`
	PublishedAt     time.Time   `json:"published_at"`
	Assets          []struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body       string `json:"body"`
}

func fetchLastRelease() (*GithubLatestRelease, error) {
	resp, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	var cResp *GithubLatestRelease = new(GithubLatestRelease)
	if err := json.NewDecoder(resp.Body).Decode(cResp); err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	_, _ = emoji.Println("Downloaded information about the last mm2 release: :white_check_mark:")
	return cResp, nil
}

func retrieveReleaseInfos(latestRelease *GithubLatestRelease, target string) (string, string, error) {
	for _, v := range latestRelease.Assets {
		if strings.Contains(v.BrowserDownloadURL, target) {
			return latestRelease.TargetCommitish, v.BrowserDownloadURL, nil
		}
	}
	return "", "", errors.New("download url not found")
}

// (hash, url, err)
func GetUrlLastMM2() (string, string, error) {
	release, err := fetchLastRelease()
	if err != nil {
		return "", "", err
	}
	switch runtime.GOOS {
	case "linux":
		return retrieveReleaseInfos(release, "Linux-Release.zip")
	case "darwin":
		return retrieveReleaseInfos(release, "Darwin-Release.zip")
	default:
		fmt.Println("Os not supported")
		return "", "", errors.New("os not supported")
	}
}
