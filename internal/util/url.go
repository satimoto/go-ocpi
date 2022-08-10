package util

import (
	"net/url"
	"strings"
)

func AppendPath(u *url.URL, path string) {
	if !strings.HasSuffix(u.Path, "/") && !strings.HasPrefix(path, "/") {
		u.Path += "/"
	}

	u.Path += path
}