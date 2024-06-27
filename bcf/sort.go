package bcf

import (
	"os/exec"
)

func OptExecSetTempDir(path string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--temp-dir", path}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// Sort cmd
func (s *Bcf) SortCmd(opts ...OptExec) *exec.Cmd {
	argsQuery := []string{"sort"}
	cmdQuery := exec.Command(s.PathBcf, argsQuery...)
	return AppendOptExec(cmdQuery, opts...)
}
