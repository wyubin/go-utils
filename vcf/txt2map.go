package vcf

import (
	"strings"

	"ailab.com/vcfgo/utils/e"
)

func Tsv2map(line string, colnames []string) (map[string]string, error) {
	res := map[string]string{}
	row := strings.SplitN(line, "\t", -1)
	rowLen := len(row)
	for idx, colname := range colnames {
		if idx >= rowLen {
			return res, e.ErrSliceIndexNotExist
		}
		valRow := row[idx]
		if valRow == "." {
			continue
		}
		res[colname] = valRow
	}
	return res, nil
}
