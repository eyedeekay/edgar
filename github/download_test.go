package github

import (
	"testing"
)

func TestReleases(t *testing.T) {
	releases, err := Releases("eyedeekay", "i2p.firefox", "eyedeekay", "")
	if err != nil {
		t.Error(err)
	}
	if len(releases) == 0 {
		t.Error("No releases found")
	}
	for i, release := range releases {
		t.Log(i, ": ", release)
	}
	t.Log("Latest Release")
	t.Log(releases[0])
}

func TestLatestRelease(t *testing.T) {
	releases, err := LatestRelease("eyedeekay", "i2p.firefox", "eyedeekay", "")
	if err != nil {
		t.Error(err)
	}
	t.Log("Latest Release")
	t.Log(releases)
}

func TestDownloadReleaseAssets(t *testing.T) {
	release, err := LatestRelease("eyedeekay", "i2p.firefox", "eyedeekay", "")
	if err != nil {
		t.Error(err)
	}
	assets, err := DownloadReleaseAssets(release)
	if err != nil {
		t.Error(err)
	}
	for i, asset := range assets {
		t.Log(i, ": ", asset)
	}
}
