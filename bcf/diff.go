package bcf

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"ailab.com/vcfgo/utils/maptool"
)

var (
	diffTidyTarget = map[string]string{
		"queryOnly":   "0000.vcf.gz",
		"dbOnly":      "0001.vcf.gz",
		"sameInQuery": "0002.vcf.gz",
		"sameInDB":    "0003.vcf.gz",
		"sites":       "sites.txt",
	}
)

// Find unique/same variants in query and db
// tidy output with queryOnly, dbOnly, sameInQuery, sameInDB, sites
func (s *Bcf) Diff(pathQuery, pathDB string, dirOutput string, argsMeta ...map[string]string) error {
	var err error
	// check pathQuery, pathDB exist
	for _, path := range []string{pathQuery, pathDB} {
		if _, err = os.Stat(path); err != nil {
			return err
		}
	}

	argsQuery := []string{"isec", "--threads", strconv.Itoa(int(s.Threads)), "-p", dirOutput, "-Oz", pathQuery, pathDB}
	cmdQuery := exec.Command(s.PathBcf, argsQuery...)
	cmdQuery.Stderr = os.Stderr
	err = AppendOptExec(cmdQuery).Run()
	if err != nil {
		return err
	}
	// tidy files
	argsMerge := map[string]string{}
	maptool.Update(argsMerge, argsMeta...)
	if len(argsMerge) == 0 {
		return nil
	}
	hasTidy := false
	// tidy and write to stderr
	for _, key := range []string{"queryOnly", "dbOnly", "sameInQuery", "sameInDB", "sites"} {
		pathTmp, found := argsMerge[key]
		if !found {
			continue
		}
		pathSrc := fmt.Sprintf("%s/%s", dirOutput, diffTidyTarget[key])
		err = os.Rename(pathSrc, pathTmp)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "tidy %s to %s\n", key, pathTmp)
		hasTidy = true
	}
	if !hasTidy {
		return nil
	}
	// move dirOutput
	fmt.Fprint(os.Stderr, "remove dirOutput\n")
	return os.RemoveAll(dirOutput)
}
