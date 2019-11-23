// +build integration

package integration

import (
	"flag"
	"log"
	"os"
	"testing"
)

var (
	// parallelism is a flag for controlling the number of linting Go routines
	// used by TestCorpus.
	parallelism = flag.Int("parallelism", 5, "number of linting Go routines to spawn")
	// configFile is a flag for specifying the config file JSON.
	configFile = flag.String("config", "./config.json", "integration test config file")
	// force is a flag for forcing the download of data files even if they are in
	// the cache dir already.
	force = flag.Bool("force", false, "ignore cached data and force new download")
	// serialSummarize is a flag for controlling whether a summary of the serial numbers
	// with lint findings (e.g. one or more fatal, error, warning or info level
	// findings) should be printed at the end of TestCorpus. Defaults to false
	// because it is very spammy with a large corpus.
	serialSummarize = flag.Bool("serialSummary", false, "print summary of all serials with lint findings")
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
)

// TestMain loads the integration test config, validates it, and prepares the
// cache (downloading configured CSV data files if needed), and then runs all tests.
func TestMain(m *testing.M) {
	flag.Parse()

	// Load and validate configuration
	c, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("error loading config file %q: %v", *configFile, err)
	}
	if err := c.Valid(); err != nil {
		log.Fatalf("error processing config file %q: %v", *configFile, err)
	}

	// Prepare cache, downloading data files if required
	if err := c.PrepareCache(*force); err != nil {
		log.Fatalf("error preparing cache: %v\n", err)
	}
	// Save the config to a global accessible to tests.
	conf = c

	// Run all tests.
	os.Exit(m.Run())
}
