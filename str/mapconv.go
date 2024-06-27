package str

import (
	"fmt"
	"strings"
)

// split string to a map
func MapConv(str string, sep string, keys []string) (map[string]string, error) {
	sliceStr := strings.Split(str, sep)
	res := map[string]string{}
	if len(sliceStr) != len(keys) {
		return nil, fmt.Errorf("length of sliceStr[%d] != length of keys[%d]", len(sliceStr), len(keys))
	}
	for idx, key := range keys {
		res[key] = sliceStr[idx]
	}
	return res, nil
}
