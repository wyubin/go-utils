package variant

import (
	"fmt"
	"regexp"

	"ailab.com/vcfgo/utils/re/base"
)

var (
	ReVariantPos = regexp.MustCompile(`^(c?h?r?[0-9XYMT]+)-([[:digit:]]+)-([ATCG]+)-([ATCG]+)$`)
)

func ValidVariants(variants []string) []string {
	return base.Valids(ReVariantPos, variants...)
}

func ResolvePos(variant string) ([]string, error) {
	strMatch := ReVariantPos.FindStringSubmatch(variant)
	if len(strMatch) == 0 {
		return nil, fmt.Errorf("invalid variant: %s", variant)
	}
	return strMatch[1:], nil
}
