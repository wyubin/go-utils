package bcf

import (
	"os/exec"
	"strconv"
)

// rename chromosome of vcf
func (s *Bcf) ReChr(pathSrc, pathRename, pathOut string) error {
	fmtOut := "v"
	if s.Compress {
		fmtOut = "z"
	}
	cmdRename := exec.Command(s.PathBcf, "annotate", "--output-type", fmtOut, "--rename-chrs", pathRename,
		"--threads", strconv.Itoa(int(s.Threads)), "-o", pathOut, pathSrc)
	return RunWithLog(cmdRename)
}
