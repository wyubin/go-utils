package bcf

import (
	"os/exec"
	"strconv"
)

// Sort cmd
func (s *Bcf) ReheaderCmd(opts ...OptExec) *exec.Cmd {
	argsQuery := []string{"reheader", "--threads", strconv.Itoa(s.Threads)}
	cmdQuery := exec.Command(s.PathBcf, argsQuery...)
	return AppendOptExec(cmdQuery, opts...)
}
