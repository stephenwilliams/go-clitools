package github

import (
	"fmt"
	"net/http"
)

func GetReleaseByTag(owner, repo, tag string) (*Release, error) {
	req, err := newRequest(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases/tags/%s", owner, repo, tag), nil)
	if err != nil {
		return nil, err
	}

	release := &Release{}
	if _, err = do(req, true, release); err != nil {
		return nil, err
	}

	return release, nil
}

func GetLatestRelease(owner, repo string) (*Release, error) {
	req, err := newRequest(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases/latest", owner, repo), nil)
	if err != nil {
		return nil, err
	}

	release := &Release{}
	if _, err = do(req, true, release); err != nil {
		return nil, err
	}

	return release, nil
}

type ReleaseSelectorFunc func([]*Release) (*Release, error)

func FindRelease(owner, repo string, selectorFunc ReleaseSelectorFunc) (*Release, error) {
	req, err := newRequest(http.MethodGet, fmt.Sprintf("/repos/%s/%s/releases", owner, repo), nil)
	if err != nil {
		return nil, err
	}

	req.URL.Query().Set("per_page", "100")

	for req != nil {
		var releases []*Release
		resp, err := do(req, true, &releases)
		if err != nil {
			return nil, err
		}

		result, err := selectorFunc(releases)
		if err != nil {
			return nil, err
		} else if result != nil {
			return result, nil
		}

		if page, err := getNextPage(resp); err != nil {
			return nil, err
		} else if page != "" {
			req, err = newRequest(http.MethodGet, page, nil)
			if err != nil {
				return nil, err
			}
		} else {
			req = nil
		}
	}

	return nil, nil
}
