package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/stephenwilliams/go-clitools/internal/iohelpers"

	"github.com/kyoh86/xdg"
)

var (
	CacheDirectory = filepath.Join(xdg.CacheHome(), "go-clitools", "cache")
	BinDirectory   = filepath.Join(xdg.DataHome(), "go-clitools", "bin")
	CacheTimeout   = time.Hour * 24 * 7

	cacheCleanedLock sync.RWMutex
	cacheCleaned     bool
)

func ensureDirectory(p string) error {
	if info, err := os.Stat(p); os.IsNotExist(err) {
		return os.MkdirAll(p, 0700)
	} else if err != nil {
		return err
	} else if !info.IsDir() {
		return fmt.Errorf("'%s' exists but is not a directory", p)
	}

	return nil
}

func cleanCacheDirectory() error {
	cacheCleanedLock.RLock()
	if cacheCleaned {
		cacheCleanedLock.RUnlock()
		return nil
	}
	cacheCleanedLock.RUnlock()

	cacheCleanedLock.Lock()
	defer cacheCleanedLock.Unlock()

	if err := filepath.Walk(CacheDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func setCache(downloader, key, version string, data interface{}) error {
	if err := ensureDirectory(CacheDirectory); err != nil {
		return err
	}

	dest := filepath.Join(CacheDirectory, fmt.Sprintf("%s__%s__%s", downloader, key, version))

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getCache(downloader, key, version string, data interface{}) (bool, error) {
	if !iohelpers.DirExists(CacheDirectory) {
		if err := os.MkdirAll(CacheDirectory, 0700); err != nil {
			return false, fmt.Errorf("failed to create cache directory: %w", err)
		}
	}

	dest := filepath.Join(CacheDirectory, fmt.Sprintf("%s__%s__%s", downloader, key, version))

	if info, err := os.Stat(dest); os.IsNotExist(err) {
		return false, nil
	} else if info.ModTime().Add(CacheTimeout).Before(time.Now()) {
		return false, nil
	}

	if err := ensureDirectory(CacheDirectory); err != nil {
		return false, err
	}

	b, err := ioutil.ReadFile(dest)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		return false, err
	}

	return true, nil
}
