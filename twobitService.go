package twobit

import (
	"bytes"
	"io/ioutil"
	"log"
)

// Service includes all services
type Service interface {
	// Genomic Interval-based services
	GenomicInterval(chr string, start, end int) (string, error)                                                                             // returns genomic sequence using 2-bit (from memory)
	GetGenomicIntervalWindow(chr string, start, end int, ref, alt string, windowL, windowR int) (seqRef, seqAlt string, err error)          // returns ref and alt with window to left and right
	GetGenomicIntervalBound(chr string, start, end int, ref, alt string, intervalStart, intervalEnd int) (seqRef, seqAlt string, err error) // returns ref and alt bound by interval (ex. interval = exon, will return exonRef and exonAlt sequences)
}

type twobitResults struct {
	tb *Reader
}

// NewDataService - Open 2-bit genome reference
func NewDataService(twobitFile string) (Service, error) {
	service := new(twobitResults)

	rdr, err := ioutil.ReadFile(twobitFile)
	if err != nil {
		log.Fatal(err)
	}

	referenceFile := bytes.NewReader(rdr)

	service.tb, err = NewReader(referenceFile)
	if err != nil {
		log.Fatal(err)
	}

	return service, nil
}
