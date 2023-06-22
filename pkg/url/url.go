package url

import "strings"

func ProtocolCheck(urlList []string) []string {
	var modifiedUrls []string

	for _, url := range urlList {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		modifiedUrls = append(modifiedUrls, url)
	}

	return modifiedUrls
}
