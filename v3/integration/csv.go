// +build integration

package integration

import (
	"encoding/base64"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"

	"github.com/zmap/zcrypto/x509"
)

// csvFieldIndex represents an index into a CSV Record.
type csvFieldIndex int

const (
	// csvSubjectDN is the index for the Subject DN CSV field.
	csvSubjectDN csvFieldIndex = iota
	// csvIssuerDN is the index for the Issued DN CSV field.
	csvIssuerDN
	// csvRaw is the index for the raw base64 encoded certificate DER CSV field.
	csvRaw
	// csvFingerprint is the index for the certificate fingerprint CSV field.
	csvFingerprint
	// end is a marker used to calculate number of fields in the CSV reader.
	end
)

// workItem is a struct collecting together a fingerprint and a parsed
// certificate that were read from a CSV record in a data file.
type workItem struct {
	// Fingerprint is the SHA256 hash of the raw certificate DER. It is provided
	// in the CSV so we capture it into a work item to avoid having to rehash the
	// DER later on.
	Fingerprint string
	// Certificate is the parsed x509 Certificate created from the CSV record's
	// Base64 encoded raw DER.
	Certificate *x509.Certificate
}

// loadCSV processes the configured data files with the provided cache
// directory, writing work items to the workChannel as they are available.
//
// Expected CSV format:
//   subject_dn, issuer_dn, raw, fingerprint_sha256
func loadCSV(workChannel chan<- workItem, directory string) {
	log.Printf("Reading data from %d CSV files", len(conf.Files))
	for i, dataFile := range conf.Files {
		path := path.Join(conf.CacheDir, dataFile.Name)
		log.Printf("Reading data from %q (%d of %d)\n",
			path, i+1, len(conf.Files))
		if err := loadCSVFile(workChannel, path, i == 0); err != nil {
			log.Fatalf("Failed reading CSV file %q: %v", path, err)
		}
		log.Printf("Done reading CSV file %q", path)
	}

	log.Printf("Finished reading data from %d CSV files. Closing work channel",
		len(conf.Files))
	close(workChannel)
}

// loadCSVFile reads and parses a certificate and fingerprint from the csvRaw
// index of each record in the provided CSV file, putting a matching work item
// into the workChannel.
func loadCSVFile(workChannel chan<- workItem, path string, skipHeader bool) error {
	// Open the input file and create a CSV reader configured for the expected
	// number of record fields.
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	in := csv.NewReader(f)
	in.FieldsPerRecord = int(end)
	in.ReuseRecord = true

	// Start reading records until there are none left.
	var skippedFirst bool
	for {
		record, err := in.Read()
		// If we read EOF its time to end the loop and return nil
		if err == io.EOF {
			return nil
		} else if err != nil {
			// If there was an error, end the loop and return non-nil
			return err
		}

		// If we haven't skipped a header yet and are configured to do so then skip
		// this record.
		if !skippedFirst && skipHeader {
			skippedFirst = true
			continue
		}

		// If a fingerprint filter is configured only include records with
		// a fingerprint that matches the filter regexp.
		if fpFilter != nil && !fpFilter.MatchString(record[csvFingerprint]) {
			continue
		}

		// Parse a certificate from the record's csvRaw index and write it to the
		// work channel.
		cert, err := parseCertificate(record[csvRaw])
		if err != nil {
			log.Printf("Warning: failed to parse record in %q: subjectDN %q fingerprint %q raw %q: %v",
				path, record[csvSubjectDN], record[csvFingerprint], record[csvRaw], err)
			continue
		}
		workChannel <- workItem{
			Fingerprint: record[csvFingerprint],
			Certificate: cert,
		}
	}
	// Control should never reach this point...
	return nil
}

// parseCertificate parses an *x509.Certificate instance from the given csvRaw
// string assumed to be the BASE64 encoding of a DER encoded x509 certificate.
func parseCertificate(csvRaw string) (*x509.Certificate, error) {
	derBytes, err := base64.StdEncoding.DecodeString(csvRaw)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
