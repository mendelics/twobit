package twobit

import (
	"fmt"
)

func (service *twobitResults) GetGenomicIntervalWindow(chr string, start, end int, ref, alt string, windowL, windowR int) (seqRef, seqAlt string, err error) {
	seqL, err := service.GenomicInterval(chr, start-windowL, start)
	if err != nil {
		return "", "", err
	}

	seqR, err := service.GenomicInterval(chr, end, end+windowR)
	if err != nil {
		return "", "", err
	}

	seqRef = fmt.Sprintf("%s%s%s", seqL, ref, seqR)
	seqAlt = fmt.Sprintf("%s%s%s", seqL, alt, seqR)

	return seqRef, seqAlt, nil
}
