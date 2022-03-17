package archives

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unzip(archive string, dest string) error {
	r, err := zip.OpenReader(archive)

	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		finfo := f.FileInfo()

		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s is an illegal filepath", fpath)
		}

		if finfo.IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		written, err := io.Copy(file, rc)
		if closeErr := file.Close(); closeErr != nil {
			rc.Close()
			return closeErr
		} else if closeErr := rc.Close(); closeErr != nil {
			return closeErr
		} else if err != nil {
			return err
		} else if written != finfo.Size() {
			return fmt.Errorf("wrote %d, want %d", written, finfo.Size())
		}
	}

	return nil
}
