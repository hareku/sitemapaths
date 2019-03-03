package urlpath

import (
	"strings"

	"golang.org/x/xerrors"
)

// GetPath returns a path from URL.
func GetPath(url string) (string, error) {
	split := strings.SplitN(url, "/", 4)

	if len(split) < 3 {
		return "", xerrors.Errorf("url is invalid: %q", url)
	}

	if len(split) == 3 {
		return "/", nil
	}

	if strings.HasPrefix(split[3], "/") {
		return split[3], nil
	}

	return "/" + split[3], nil
}
