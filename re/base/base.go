package base

import "regexp"

func Valids(reObj *regexp.Regexp, strSlice ...string) []string {
	valids := []string{}
	for _, str := range strSlice {
		if reObj.MatchString(str) {
			valids = append(valids, str)
		}
	}
	return valids
}
