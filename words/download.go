// gather a wordlist from a web source

package words

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const (
	SourceAlpha  = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"
	SourceMedium = "https://raw.githubusercontent.com/zekenaulty/" +
		"Dataset-google-10000-english/refs/heads/master/google-10000-english-usa-no-swears-medium.txt"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CacheWords downloads a file from the specified URL and saves it to the given directory.
// If the file already exists in the directory, it skips the download and returns the file name.
//
// Parameters:
//   - site: The URL of the file to download.
//   - directory: The directory where the file should be saved.
//
// Returns:
//   - A string representing the path to the cached file.
//   - An error if any issues occur during URL parsing, downloading, file creation, or writing.
//
// The function checks if the file already exists in the specified directory
// before attempting to download it. If the file does not exist, it downloads
// the content from the URL, saves it to the directory, and returns the file path.
func CacheWords(site string, directory string) (string, error) {
	u, err := url.Parse(site)
	if err != nil {
		return "", fmt.Errorf("parsing URL: %w", err)
	}

	basename := filepath.Base(u.Path)
	if fileExists(basename) {
		return basename, nil
	}

	resp, err := http.Get(site)
	if err != nil {
		return "", fmt.Errorf("downloading content: %w", err)
	}
	defer resp.Body.Close()

	output, err := os.Create(filepath.Join(directory, basename))
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer output.Close()

	if _, err := io.Copy(output, resp.Body); err != nil {
		return "", fmt.Errorf("writing content: %w", err)
	}

	return output.Name(), nil
}
