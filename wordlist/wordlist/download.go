// gather a wordlist from github

package wordlist

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
//   - filename: The name of the file where the content is saved.
//
// Note: The function assumes the existence of helper functions `check` and `fileExists`
// for error handling and file existence checking, respectively.
func cacheWords(site string) (filename string) {
	u, err := url.Parse(site)
	check(err)
	filename = filepath.Base(u.Path)
	fmt.Println(filename)

	if fileExists(filename) {
		log.Printf("File %s already exists, skipping download.", filename)
		return filename
	}

	resp, err := http.Get(site)
	check(err)
	defer resp.Body.Close()

	// write the word list to a file
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	check(err)

	return filename
}
