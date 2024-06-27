package str

import "regexp"

const (
	TimeFmtWayBack = "20060102150405"
)

var (
	ReWayBack = regexp.MustCompile(`\d{14}`)
)

func SortFuncWayBack(idxReverse bool) func(a, b string) int {
	numMul := 1
	if idxReverse {
		numMul = -1
	}
	return func(a, b string) int {
		aWayBack := ReWayBack.FindString(a)
		bWayBack := ReWayBack.FindString(b)
		if aWayBack > bWayBack {
			return numMul
		}
		if aWayBack < bWayBack {
			return -numMul
		}
		return 0
	}
}
