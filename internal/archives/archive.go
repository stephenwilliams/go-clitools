package archives

import (
	"fmt"
	"os"
	"strings"
)

func Unarchive(archive, dest string) error {
	if info, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if !info.IsDir() {
		return fmt.Errorf("dir is not a directory '%s'", dest)
	}

	if strings.HasSuffix(archive, ".tar.gz") ||
		strings.HasSuffix(archive, ".tar") ||
		strings.HasSuffix(archive, ".tgz") {
		return untar(archive, dest)
	} else if strings.HasSuffix(archive, ".zip") {
		return unzip(archive, dest)
	} else {
		return fmt.Errorf("unknown archive type for '%s'", archive)
	}
}
