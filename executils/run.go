package executils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type OptExec func(execInput *exec.Cmd)

func appendOptExec(execInput *exec.Cmd, opts ...OptExec) *exec.Cmd {
	for _, opt := range opts {
		opt(execInput)
	}
	return execInput
}

// Run with log
func RunWithLog(execInput *exec.Cmd, opts ...OptExec) error {
	var errb bytes.Buffer
	execInput.Stderr = &errb
	appendOptExec(execInput, opts...)
	err := execInput.Run()
	if err != nil {
		errb.WriteTo(os.Stderr)
		fmt.Printf("Error cmd: %s %s\n", execInput.Path, strings.Join(execInput.Args, " "))
	}
	return err
}
