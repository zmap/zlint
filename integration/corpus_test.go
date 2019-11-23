// +build integration

package integration

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/zmap/zlint"
	"github.com/zmap/zlint/lints"
)

// lintCertificate lints the provided work item's certificate to produce
// a certResult that can be used to determine which lint results the certificate
// had without maintaining the full ResultSet.
func lintCertificate(work workItem) certResult {
	// Lint the certiifcate to produce a full result set
	result := certResult{
		Fingerprint: work.Fingerprint,
	}
	resultSet := zlint.LintCertificate(work.Certificate)
	for _, r := range resultSet.Results {
		switch r.Status {
		case lints.Notice:
			result.Result.NoticeCount++
		case lints.Warn:
			result.Result.WarnCount++
		case lints.Error:
			result.Result.ErrCount++
		case lints.Fatal:
			result.Result.FatalCount++
		}
	}
	return result
}

// TestCorpus concurrently reads certificates from each of the global conf's CSV
// data files while in parallel linting the certificates and counting how many
// of each lint result are produced across all data files. The lint result
// totals are enforced against the expected values from the global conf.
func TestCorpus(t *testing.T) {
	// Create a work channel with enough capacity to let each loader write
	// 1 work item without blocking.
	workChannel := make(chan workItem, len(conf.Files))

	// Start loading certificates from the config CSV files. This is done in
	// a separate Go routine because loadCSV will block until completion. We want
	// to let the test continue to run so certificates can be linted as they
	// arrive.
	go func() {
		loadCSV(workChannel, conf.CacheDir)
	}()

	log.Printf(
		"Linting certificates using %d Go routines. "+
			"Printing one '.' per %d certificates",
		*parallelism, *outputTick)

	// Create *parallelism separate Go routines for reading certificates from
	// the work channel, linting them, and writing the result to a results
	// channel.
	results := make(chan certResult, *parallelism)
	var wg sync.WaitGroup
	for i := 0; i < *parallelism; i++ {
		wg.Add(1)
		go func() {
			// Read work until the channel is closed
			for c := range workChannel {
				results <- lintCertificate(c)
			}
			// Once the workChannel has closed this routine is done.
			wg.Done()
		}()
	}

	// Also start a Go routine to read from the results channel, aggregating the
	// results into the results map
	var total int
	var fatalResults int
	resultsMap := make(map[string]result)
	doneChan := make(chan bool, 1)
	go func() {
		// Read results as they arrive on the channel until it is closed.
		for r := range results {
			// Count fatal results separately since this should always be 0
			fatalResults += int(r.Result.FatalCount)
			// if the result had some error/warn/info findings, track the fingerprint
			// in the results map.
			if !r.Result.fullPass() {
				resultsMap[r.Fingerprint] = r.Result
			}
			// Every *outputTick certificate results print a '.' to keep CI from thinking this
			// long running job is dead in the water.
			total++
			if total%*outputTick == 0 {
				fmt.Printf(".")
			}
		}
		// Once the results channel is closed and we're done tabulating in this
		// routine write to the doneChan so the test can complete.
		doneChan <- true
	}()

	// Wait for the work channel to be drained by all of the workers.
	wg.Wait()
	// Close the results channel
	close(results)
	// Wait for the results tabulation routine to complete.
	<-doneChan

	// Verify results match the conf's expected totals.
	t.Logf("linted %d certificates", total)
	// There should never be any fatal results.
	if fatalResults != 0 {
		t.Errorf("expected 0 fatal results, found %d\n", fatalResults)
	}

	if *serialSummarize {
		for serial, result := range resultsMap {
			fmt.Printf("%s\t%s\n", serial, result)
		}
	}

	// No expected to confirm against, save a new expected
	if len(conf.Expected) == 0 {
		t.Logf("config file %q had no expected map to enforce results against",
			*configFile)
	} else {
		// Otherwise enforce the maps match
		for k, v := range resultsMap {
			if conf.Expected[k] != v {
				t.Errorf("expected serial %q to have result %s got %s\n",
					k, conf.Expected[k], v)
			}
		}
	}

	// If *overwriteExpected is true overwrite the expected map with the results
	// from this run and save the updated configuration to disk. If there were
	// t.Errorf's in this run then they will pass next run because the
	// expectations will match reality. This should primarily be used to bootstrap
	// an initial expectedMap or to update the expectedMap with vetted changes to
	// the corpus that result from new lints, bugfixes, etc.
	if *overwriteExpected {
		t.Logf("overwriting expected map in config file %q",
			*configFile)
		conf.Expected = resultsMap
		if err := conf.Save(*configFile); err != nil {
			t.Errorf("failed to save expected map to config file %q: %v", *configFile, err)
		}
	}
}
