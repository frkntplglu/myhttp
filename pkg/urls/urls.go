package urls

import (
	"net/url"
	"strings"
)

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

func IsValidUrl(u string) bool {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false
	}

	return parsedUrl.Scheme != "" && parsedUrl.Host != ""
}
