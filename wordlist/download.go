// gather a wordlist from github

package wordlist

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// cacheWords downloads the content from the given URL and saves it to a file.
// If the file already exists, it skips the download and returns the filename.
//
// Parameters:
//   - site: The URL of the site to download the content from.
//
// Returns:
func cacheWords(site string) (string, error) {
	u, err := url.Parse(site)
	if err != nil {
		return "", fmt.Errorf("parsing URL: %w", err)
	}

	filename := filepath.Base(u.Path)
	if fileExists(filename) {
		return filename, nil
	}

	resp, err := http.Get(site)
	if err != nil {
		return "", fmt.Errorf("downloading content: %w", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return "", fmt.Errorf("writing content: %w", err)
	}

	return filename, nil
}
