package bcf

import (
	"os/exec"
	"strconv"
)

// split multiallelic
func OptExecNormMultialle() OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-m", "-any", "--old-rec-tag", "oldMultiallelic"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecNormRef(pathRef string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-f", pathRef, "-c", "s"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecNormRemoveDuplicate() OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-d", "none"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func (s *Bcf) NormCmd(opts ...OptExec) *exec.Cmd {
	cmdSlice := []string{"norm", "--threads", strconv.Itoa(int(s.Threads))}
	cmdNorm := exec.Command(s.PathBcf, cmdSlice...)
	return AppendOptExec(cmdNorm, opts...)
}
