package ver

import (
	"fmt"
	"os"
	"regexp"

	"golang.org/x/mod/semver"
)

var (
	ReVer = regexp.MustCompile(`v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)
)

// get version map from parsing []string
func VersionMap(strItems []string) map[string]string {
	res := map[string]string{}
	for _, name := range strItems {
		strMatch := ReVer.FindString(name)
		if len(strMatch) == 0 {
			continue
		}
		res[strMatch] = name
	}
	return res
}

func UseVersion(strItems []string, version string) (string, error) {
	return selVersion(VersionMap(strItems), version)
}

// use version from Dir
func UseVersionDir(dir, version string) (string, error) {
	items := []string{}
	// get versions map
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("dir not found: %s", dir)
	}
	for _, file := range files {
		items = append(items, file.Name())
	}
	return UseVersion(items, version)
}

func selVersion(verMap map[string]string, version string) (string, error) {
	if ReVer.MatchString(version) {
		resTmp, found := verMap[version]
		if !found {
			return "", fmt.Errorf("version not found: %s", version)
		}
		return resTmp, nil
	} else {
		switch version {
		case "latest":
			return getLastVersion(verMap), nil
		default:
			return "", fmt.Errorf("invalid version: %s", version)
		}
	}
}

func getLastVersion(verMap map[string]string) string {
	vers := []string{}
	for ver := range verMap {
		vers = append(vers, ver)
	}
	semver.Sort(vers)
	return verMap[vers[len(vers)-1]]
}
