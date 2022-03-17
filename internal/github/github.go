package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	HTTP = http.Client{
		Timeout: time.Second * 30,
	}
	ApiEndpoint *url.URL
	UserAgent   = "go-clitools/1.0"
	Token       = os.Getenv("GITHUB_TOKEN")

	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("received bad request")
	ErrInternalServerError = errors.New("internal server error")
)

func init() {
	var err error
	ApiEndpoint, err = url.Parse("https://api.github.com")
	if err != nil {
		panic(err)
	}
}

// based off of the link header parsing from go-github
// https://github.com/google/go-github/blob/c550a55450088b424c856d9493a342953de9faf8/github/github.go#L456
func getNextPage(r *http.Response) (string, error) {
	links := r.Header.Values("link")
	if len(links) == 0 {
		return "", nil
	}

	for _, link := range strings.Split(links[0], ",") {
		segments := strings.Split(strings.TrimSpace(link), ";")

		// link must at least have href and rel
		if len(segments) < 2 {
			continue
		}

		// ensure href is properly formatted
		if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
			continue
		}

		url, err := url.Parse(segments[0][1 : len(segments[0])-1])
		if err != nil {
			continue
		}

		for _, segment := range segments[1:] {
			if strings.TrimSpace(segment) == "rel=\"next\"" {
				return url.String(), nil
			}
		}
	}

	return "", nil
}

func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	part, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	relURL := ApiEndpoint.ResolveReference(part)

	req, err := http.NewRequest(method, relURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed creating request: %w", err)
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	if Token != "" {
		req.Header.Set("Authentication", fmt.Sprintf("token %s", Token))
	}

	return req, nil
}

func do(req *http.Request, readValue bool, v interface{}) (*http.Response, error) {
	resp, err := HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return resp, ErrBadRequest
	case http.StatusInternalServerError:
		return resp, ErrInternalServerError
	case http.StatusNotFound:
		return resp, ErrNotFound
	}

	if !readValue {
		return resp, nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
