package genomedata

import (
	"bufio"
	"fmt"
	"io"
)

func ReadFirstLine(r io.Reader) (string, error) {
	var res string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Get first line to check chr and exit
		res = scanner.Text()
		break
	}
	if res == "" {
		return res, fmt.Errorf("file reach EOF")
	}
	return res, nil
}
