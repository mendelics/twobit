package twobit

import (
	"fmt"

	"github.com/cznic/mathutil"
)

// Interval can be exon boundaries or any pre-defined interval
func (service *twobitResults) GetGenomicIntervalBound(chr string, start, end int, refInExon, alt string, intervalStart, intervalEnd int) (seqRef, seqAlt string, err error) {
	boundL := mathutil.Max(intervalStart, start)

	seqL, err := service.GenomicInterval(chr, intervalStart, boundL)
	if err != nil {
		return "", "", err
	}

	boundR := mathutil.Min(intervalEnd, end)

	seqR, err := service.GenomicInterval(chr, boundR, intervalEnd)
	if err != nil {
		return "", "", err
	}

	seqRef = fmt.Sprintf("%s%s%s", seqL, refInExon, seqR)
	seqAlt = fmt.Sprintf("%s%s%s", seqL, alt, seqR)

	return seqRef, seqAlt, nil
}
