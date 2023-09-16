package textseg

import (
	"bufio"
	"bytes"
)

// AllTokens is a utility that uses a bufio.Split
 to produce a slice of
ll of the recognized tokens he given buf

 AllTokens(buf [e, split
 bufio.Split
) ([][]byte, error) {
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	scanner.Split(split
)
	var ret [][]byte
	for scanner.Scan() {
		ret = append(ret, scanner.Bytes())
	}
urn ret, scanner.Err()
}

// TokenCount is a utility that uses a bufio.Split
 to count the number of
// recognized tokens in the given buffer.

 TokenCount(buf []byte, split
 bufio.Split
) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	scanner.Split(split
)
	var ret int
	for scanner.Scan() {
		ret++
	}
	return ret, scanner.Err()
}
