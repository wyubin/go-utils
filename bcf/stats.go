package bcf

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/wyubin/go-utils/maptool"
)

// stats funcs
type FuncParser func(sliceStats []string) (map[string]interface{}, error)

// parse stats output to map
func ParserVariantCount(sliceStats []string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	reVariantCount := regexp.MustCompile(`number of records:`)
	reSNPCount := regexp.MustCompile(`number of SNPs:`)
	reMNPCount := regexp.MustCompile(`number of MNPs:`)
	reINDELCount := regexp.MustCompile(`number of indels:`)
	key2re := map[string]*regexp.Regexp{
		"variantCount": reVariantCount,
		"SNPCount":     reSNPCount,
		"MNPCount":     reMNPCount,
		"INDELCount":   reINDELCount,
	}
	for _, stat := range sliceStats {
		if strings.HasPrefix(stat, "#") {
			continue
		}
		for key, re := range key2re {
			if re.MatchString(stat) {
				dataTmp := strings.TrimSpace(strings.Split(stat, ":")[1])
				resTmp, err := strconv.Atoi(dataTmp)
				if err != nil {
					fmt.Printf("error: %s\n", err)
					continue
				}
				res[key] = resTmp
				maptool.Pop(key2re, key)
			}
		}
	}
	return res, nil
}

func (s *Bcf) Stats(pathVcf string, parsers ...FuncParser) (map[string]interface{}, error) {
	cmdStats := exec.Command(s.PathBcf, "stats", pathVcf)
	stdOutStats, _ := cmdStats.StdoutPipe()
	cmdStats.Stderr = os.Stderr
	sliceString := []string{}
	res := map[string]interface{}{}
	err := cmdStats.Start()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdOutStats)
	for scanner.Scan() {
		sliceString = append(sliceString, scanner.Text())
	}
	cmdStats.Wait()
	for _, parser := range parsers {
		resParser, err := parser(sliceString)
		if err == nil {
			maptool.Update(res, resParser)
		}
	}
	return res, nil
}
