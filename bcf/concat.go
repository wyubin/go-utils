package bcf

import (
	"os/exec"
	"strconv"
)

// concat vcf files based on variant pos
func (s *Bcf) ConcatCmd(appendFiles []string, opts ...OptExec) *exec.Cmd {
	argSlice := []string{"concat", "-a", "-d", "exact", "--threads", strconv.Itoa(int(s.Threads))}
	argSlice = append(argSlice, appendFiles...)
	cmds := exec.Command(s.PathBcf, argSlice...)
	return AppendOptExec(cmds, opts...)
}
