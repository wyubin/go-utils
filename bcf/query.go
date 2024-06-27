package bcf

import (
	"os/exec"
)

func OptExecQueryHeader() OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--print-header"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// Query Vcf
func (s *Bcf) QueryCmd(strFmt string, opts ...OptExec) *exec.Cmd {
	argsQuery := []string{"query", "-f", strFmt}
	cmdQuery := exec.Command(s.PathBcf, argsQuery...)
	return AppendOptExec(cmdQuery, opts...)
}
