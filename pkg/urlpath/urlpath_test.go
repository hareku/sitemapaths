package urlpath

import "testing"

func TestGetPath(t *testing.T) {
	testSet := map[string]string{
		"https://example.com":                  "/",
		"https://example.com/my-post":          "/my-post",
		"https://example.com/my-post?param=1":  "/my-post?param=1",
		"https://example.com/posts/first-post": "/posts/first-post",
	}

	for url, expect := range testSet {
		actual, _ := GetPath(url)
		if actual != expect {
			t.Errorf("GetPath error: url=%q, expect=%q, actual=%q", url, expect, actual)
		}
	}

	_, err := GetPath("https:invalid-url")
	if err == nil {
		t.Error("GetPath should return an error if that got invalid URL")
	}
}
