package twobit

import (
	"errors"
	"fmt"
	"strings"
)

// GenomicInterval - Receives chr, start, end and returns genome reference.
func (service *twobitResults) GenomicInterval(chr string, start, end int) (string, error) {
	if chr == "" {
		return "", errors.New("GenomicInterval can't accept blank chromosome")
	}

	chrName := chr

	if !strings.HasPrefix(chr, "chr") {
		chrName = fmt.Sprintf("chr%s", chr)
	}

	if chrName == "chrMT" {
		chrName = "chrM"
	}

	if start == end {
		return "", nil
	}

	seq, err := service.tb.ReadRange(chrName, start, end)

	return string(seq), err
}
