// +build integration

package integration

import (
	"flag"
	"log"
	"os"
	"regexp"
	"testing"
)

var (
	// parallelism is a flag for controlling the number of linting Go routines
	// used by TestCorpus.
	parallelism = flag.Int("parallelism", 5, "number of linting Go routines to spawn")
	// configFile is a flag for specifying the config file JSON.
	configFile = flag.String("config", "./config.json", "integration test config file")
	// forceDownload is a flag for forcing the download of data files even if they are in
	// the cache dir already.
	forceDownload = flag.Bool("forceDownload", false, "ignore cached data and force new download")
	// saveExpected is a flag for controlling whether the expectedMap is saved to
	// the configuration or not.
	overwriteExpected = flag.Bool("overwriteExpected", false, "save test results as the new expected map in config file")
	// fpSummarize is a flag for controlling whether a summary of the cert fingerprints
	// with lint findings (e.g. one or more fatal, error, warning or info level
	// findings) should be printed at the end of TestCorpus. Defaults to false
	// because it is very spammy with a large corpus.
	fpSummarize = flag.Bool("fingerprintSummary", false, "print summary of all certificate fingerprints with lint findings")
	// lintSummarize is a flag for controlling whether a summary of result types
	// by lint name is printed at the end of TestCorpus. Defaults to false because
	// it is very spammy with a large corpus.
	lintSummarize = flag.Bool("lintSummary", false, "print summary of result type counts by lint name")
	// fpFilterString is a flag for controlling which certificate fingerprints are run
	// through the lints.
	fpFilterString = flag.String("fingerprintFilter", "", "if not-empty only certificate fingerprints that match the provided regexp will be run")
	// lintFilterString is a flag for controlling which lints are run against the test
	// corpus.
	lintFilterString = flag.String("lintFilter", "", "if not-empty only lints with a name that match the provided regexp will be run")
	// outputTick is a flag for controlling the number of certificates that are
	// linted before a '.' is printed in the console. This controls the mechanism
	// used to keep Travis from thinking the job is dead because there hasn't been
	// output.
	outputTick = flag.Int("outputTick", 1000,
		"number of certificates to lint before printing a '.' marker in the output")
)

var (
	// config is a global var for the integration test configuration.
	conf *config

	// fpFilter and lintFilter are regexps for filtering certificate fingerprints
	// to be linted and lints to be run.
	fpFilter, lintFilter *regexp.Regexp
)

// TestMain loads the integration test config, validates it, and prepares the
// cache (downloading configured CSV data files if needed), and then runs all tests.
func TestMain(m *testing.M) {
	flag.Parse()

	if *fpFilterString != "" {
		filter, err := regexp.Compile(*fpFilterString)
		if err != nil {
			log.Fatalf("error compiling -fingerprintFilter regexp %q: %v", *fpFilterString, err)
		}
		fpFilter = filter
	}

	if *lintFilterString != "" {
		filter, err := regexp.Compile(*lintFilterString)
		if err != nil {
			log.Fatalf("error compiling -lintFilter regexp %q: %v", *lintFilterString, err)
		}
		lintFilter = filter
	}

	// Load and validate configuration
	c, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("error loading config file %q: %v", *configFile, err)
	}
	if err := c.Valid(); err != nil {
		log.Fatalf("error processing config file %q: %v", *configFile, err)
	}

	// Prepare cache, downloading data files if required (or if forced by user
	// request with forceDownload)
	if err := c.PrepareCache(*forceDownload); err != nil {
		log.Fatalf("error preparing cache: %v\n", err)
	}
	// Save the config to a global accessible to tests.
	conf = c

	// Run all tests.
	os.Exit(m.Run())
}
