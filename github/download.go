package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/github-release/github-release/github"
)

type Release struct {
	Url         string     `json:"url"`
	PageUrl     string     `json:"html_url"`
	UploadUrl   string     `json:"upload_url"`
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"body"`
	TagName     string     `json:"tag_name"`
	Draft       bool       `json:"draft"`
	Prerelease  bool       `json:"prerelease"`
	Created     *time.Time `json:"created_at"`
	Published   *time.Time `json:"published_at"`
	Assets      []Asset    `json:"assets"`
}

type Asset struct {
	Url         string    `json:"url"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"content_type"`
	State       string    `json:"state"`
	Size        uint64    `json:"size"`
	Downloads   uint64    `json:"download_count"`
	Created     time.Time `json:"created_at"`
	Published   time.Time `json:"published_at"`
}

var EnvToken = os.Getenv("GITHUB_TOKEN")
var EnvUser = os.Getenv("GITHUB_USER")
var EnvAuthUser = os.Getenv("GITHUB_AUTH_USER")
var EnvRepo = os.Getenv("GITHUB_REPO")
var EnvApiEndpoint = os.Getenv("GITHUB_API")

const (
	RELEASE_LIST_URI    = "/repos/%s/%s/releases"
	RELEASE_LATEST_URI  = "/repos/%s/%s/releases/latest"
	RELEASE_DATE_FORMAT = "02/01/2006 at 15:04"
)

func Releases(user, repo, authUser, token string) ([]Release, error) {
	if token == "" {
		token = EnvToken
	}
	if user == "" {
		user = EnvUser
	}
	if repo == "" {
		repo = EnvRepo
	}
	if authUser == "" {
		authUser = EnvAuthUser
	}
	log.Println("releases", user, repo)
	var releases []Release
	client := github.NewClient(authUser, token, nil)
	client.SetBaseURL(EnvApiEndpoint)
	err := client.Get(fmt.Sprintf(RELEASE_LIST_URI, user, repo), &releases)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

func LatestRelease(user, repo, authUser, token string) (Release, error) {
	releases, err := Releases(user, repo, authUser, token)
	if err != nil {
		return Release{}, err
	}
	return releases[0], err
}

func downloadBytes(url, dest string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func downloadFile(url, dest string) error {
	log.Println("downloading to", dest)
	if _, err := os.Stat(dest); err == nil {
		// move the file out of the way to dest.$date.bak
		date := time.Now().Format("20060102")
		bakDest := fmt.Sprintf("%s.%s.bak", dest, date)
		err := os.Rename(dest, bakDest)
		if err != nil {
			return err
		}
	}
	jsonBytes, err := downloadBytes(url, dest)
	if err != nil {
		return err
	}
	// marshal the json into a map
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		return err
	}
	if _, ok := jsonMap["browser_download_url"]; ok {
		downloadUrl := jsonMap["browser_download_url"].(string)
		log.Println("downloading from", downloadUrl)
		resp, err := http.Get(downloadUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		out, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		return err
	}
	return nil
}

func DownloadReleaseAssets(release Release) ([]Asset, error) {
	if len(release.Assets) == 0 {
		return nil, fmt.Errorf("No assets found")
	}
	for i := range release.Assets {
		downloadFile(release.Assets[i].Url, release.Assets[i].Name)
	}
	return release.Assets, nil
}

func DownloadLatestReleaseAssets(user, repo, authUser, token string) ([]Asset, error) {
	release, err := LatestRelease(user, repo, authUser, token)
	if err != nil {
		return nil, err
	}
	return DownloadReleaseAssets(release)
}
