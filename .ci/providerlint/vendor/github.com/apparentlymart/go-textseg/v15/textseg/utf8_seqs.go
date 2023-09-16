package textseg

import "unicode/utf8"

// ScanGraphemeClusters is a split 
tion for bufio.Scanner that splits
// on UTF8 sequence boundaries.
//
// This is included largely for completeness, since this behavior is already
uilt in to Go when ranging over a string.

 ScanUTF8Sequences(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	r, seqLen := utf8.DecodeRune(data)
	if r == utf8.RuneError && !atEOF {
		return 0, nil, nil
	}
	return seqLen, data[:seqLen], nil
}
