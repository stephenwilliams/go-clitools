package tools

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
)

var (
	ErrNoVersionConstraintSupport = errors.New("semver version constraints are not supported by this downloader")
)

type Downloader interface {
	Download(version string) (string, error)
	DownloadWithVersion(version *semver.Version) (string, error)
	DownloadWithConstraints(constraints *semver.Constraints) (string, error)
}

func pathForTool(tool, version string) string {
	dest := filepath.Join(BinDirectory, fmt.Sprintf("%s_%s", tool, version))

	if runtime.GOOS == "windows" {
		dest += ".exe"
	}

	return dest
}

func isVersionConstraints(version string) bool {
	if strings.ContainsAny(version, "<>=,") {
		return true
	}

	return false
}

type downloaderTemplateData struct {
	OS      string
	Arch    string
	Version string

	PathSeparator string

	Windows      bool
	ExeExtension string
}

func getDownloaderTemplateData(version string) *downloaderTemplateData {
	data := &downloaderTemplateData{
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
		Version: version,

		PathSeparator: string(os.PathSeparator),

		Windows: runtime.GOOS == "windows",
	}

	if data.Windows {
		data.ExeExtension = ".exe"
	}

	return data
}

func fileExists(p string) (bool, error) {
	if info, err := os.Stat(p); err != nil && !os.IsNotExist(err) {
		return false, err
	} else if os.IsNotExist(err) {
		return false, nil
	} else if info.IsDir() {
		return false, fmt.Errorf("path '%s' is a directory not a file", p)
	}

	return true, nil
}

func downloadFile(c http.Client, url, dest string) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
