package vcf

import (
	"os"
	"strconv"
	"strings"
)

func GetHumanChrList(chrMode bool) []string {
	chrList := []string{}
	for chrCount := 1; chrCount < 23; {
		chrList = append(chrList, strconv.Itoa(chrCount))
		chrCount += 1
	}
	chrList = append(chrList, []string{"X", "Y", "MT"}...)
	if chrMode {
		indLast := len(chrList) - 1
		for ind, chr := range chrList[:indLast] {
			chrList[ind] = "chr" + chr
		}
		chrList[indLast] = "chrM"
	}
	return chrList
}

func DumpHumanChrList(chrMode bool, pathDump string) error {
	fChr, err := os.Create(pathDump)
	if err != nil {
		return err
	}
	defer fChr.Close()
	currentChrs := GetHumanChrList(chrMode)
	newChrs := GetHumanChrList(!chrMode)
	for idx, chr := range currentChrs {
		_, err = fChr.WriteString(strings.Join([]string{chr, newChrs[idx]}, "\t") + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
