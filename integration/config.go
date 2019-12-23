// +build integration

package integration

import (
	"compress/bzip2"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// dataFile is a struct describing a named CSV data file that can be downloaded
// from a URL when it is not present already on disk. If the URL ends in "bz2"
// then the data at the given URL is assumed to be compressed with Bzip2 and
// will be automatically decompressed when fetching the URL to write the data
// file to disk. By default the first datafile in the set is assumed to have
// a header line that must be skipped for data processing.
type dataFile struct {
	Name string
	URL  string
}

// Valid returns an error if the data file has an empty name or URL.
func (f dataFile) Valid() error {
	if f.Name == "" {
		return errors.New("Name is empty")
	}
	if f.URL == "" {
		return errors.New("URL is empty")
	}
	return nil
}

// ExistsIn checks if a file matching the data file's name exists in the
// provided directory.
func (f dataFile) ExistsIn(dir string) (bool, error) {
	p := path.Join(dir, f.Name)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// DownloadTo will fetch the data file from its URL and write the contents to
// a file in the provided directory, handling Gzip2 decompression if required.
// An error is returned if fetching the URL fails, or if the remote server
// returns a HTTP status code other than 200.
func (f dataFile) DownloadTo(dir string) error {
	p := path.Join(dir, f.Name)

	resp, err := http.Get(f.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if expected := http.StatusOK; resp.StatusCode != expected {
		return fmt.Errorf("bad HTTP response from %q: %d != %d\n",
			f.URL, resp.StatusCode, expected)
	}

	var reader io.Reader = resp.Body
	if strings.HasSuffix(f.URL, ".bz2") {
		reader = bzip2.NewReader(reader)
	}

	dataBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(p, dataBytes, 0644); err != nil {
		return err
	}

	return nil
}

// config is a struct holding integration test configuration data.
type config struct {
	CacheDir string
	Files    []dataFile
	Expected keyedCounts
}

// loadConfig returns a config struct populated from the JSON serialization in
// the given file or returns an error if reading or unmarshaling the config file
// fails.
func loadConfig(file string) (*config, error) {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var c config
	if err := json.Unmarshal(jsonBytes, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Save persists a config in JSON form to the given file or returns an error.
func (c *config) Save(file string) error {
	jsonBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, jsonBytes, 0644)
}

// Valid returns an error if the config has an empty CacheDir, no Files, or if
// any of the Files are not valid data file configs.
func (c config) Valid() error {
	if c.CacheDir == "" {
		return errors.New("no CacheDir defined")
	}
	if len(c.Files) == 0 {
		return errors.New("No Files defined")
	}
	for i, file := range c.Files {
		if err := file.Valid(); err != nil {
			return fmt.Errorf("File %d was not valid: %v\n", i, err)
		}
	}
	return nil
}

// PrepareCache creates the CacheDir if it does not exist and will download any
// of the Files that are not present in the CacheDir. If force is true then data
// files will be downloaded even if they are already present in the cachedir.
// This can be used to force an update when the upstream file content has
// changed and a stale copy exists in the cache
func (c config) PrepareCache(force bool) error {
	if _, err := os.Stat(c.CacheDir); os.IsNotExist(err) {
		log.Printf("Creating cache directory %q\n", c.CacheDir)
		os.Mkdir(c.CacheDir, 0744)
	} else {
		log.Printf("Using existing cache directory %q\n", c.CacheDir)
	}
	for i, f := range c.Files {
		if exists, err := f.ExistsIn(c.CacheDir); err != nil {
			log.Fatalf("error checking cache: %v\n", err)
		} else if !exists || force {
			log.Printf("Downloading data file %q (%d of %d, url: %q)",
				f.Name, i+1, len(c.Files), f.URL)
			if err := f.DownloadTo(c.CacheDir); err != nil {
				log.Fatalf("Failed to download: %v", err)
			}
			log.Printf("Download complete")
		} else {
			log.Printf("Using cached data file %q", f.Name)
		}
	}
	return nil
}
