package bcf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"slices"
)

type OptExec func(execInput *exec.Cmd)

func OptExecSetCompress(needCompress bool) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--output-type"}
		if needCompress {
			args = append(args, "z")
			if !slices.Contains([]string{"sort", "view", "norm"}, c.Args[1]) {
				args = append(args, "--write-index")
			}
		} else {
			args = append(args, "v")
		}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func OptExecSetInputPath(path string) OptExec {
	return func(c *exec.Cmd) {
		c.Args = append(c.Args, path)
	}
}

func OptExecSetOutputPath(path string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-o", path}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

func AppendOptExec(execInput *exec.Cmd, opts ...OptExec) *exec.Cmd {
	for _, opt := range opts {
		opt(execInput)
	}
	return execInput
}

// Run with log
func RunWithLog(execInput *exec.Cmd, opts ...OptExec) error {
	var errb bytes.Buffer
	execInput.Stderr = &errb
	AppendOptExec(execInput, opts...)
	err := execInput.Run()
	if err != nil {
		errb.WriteTo(os.Stderr)
		fmt.Printf("Error cmd: %s %s\n", execInput.Path, strings.Join(execInput.Args, " "))
	}
	return err
}
