package bcf

import (
	"os/exec"
	"strconv"
	"strings"
)

func OptExecAnnotateHeader(pathHeader string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-h", pathHeader}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecAnnotateRename(pathRename string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--rename-chrs", pathRename}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecAnnotateAppend(pathSrc string, tags []string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-a", pathSrc, "-c", strings.Join(tags, ",")}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecAnnotateRemove(tags ...string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-x", strings.Join(tags, ",")}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// build annotate cmd
func (s *Bcf) AnnotateCmd(opts ...OptExec) *exec.Cmd {
	cmdNorm := exec.Command(s.PathBcf, "annotate", "--threads", strconv.Itoa(s.Threads))
	return AppendOptExec(cmdNorm, opts...)
}
