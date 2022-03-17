package archives

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func untar(archive, dest string) error {
	tarFile, err := os.Open(archive)
	if err != nil {
		return err
	}

	defer tarFile.Close()

	var tr *tar.Reader
	if strings.HasSuffix(archive, ".gz") || strings.HasSuffix(archive, ".tgz") {
		gz, err := gzip.NewReader(tarFile)
		if err != nil {
			return err
		}
		defer gz.Close()
		tr = tar.NewReader(gz)
	} else {
		tr = tar.NewReader(tarFile)
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		finfo := header.FileInfo()
		fpath := filepath.Join(dest, header.Name)

		if finfo.Mode().IsDir() {
			if err := os.MkdirAll(fpath, 0755); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, finfo.Mode().Perm())
		if err != nil {
			return err
		}

		written, cpErr := io.Copy(file, tr)
		if closeErr := file.Close(); closeErr != nil {
			return err
		} else if cpErr != nil {
			return cpErr
		} else if written != finfo.Size() {
			return fmt.Errorf("wrote %d, want %d", written, finfo.Size())
		}
	}
	return nil
}
