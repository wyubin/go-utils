package bcf

import (
	"os/exec"
	"strconv"
	"strings"
)

// without header
func OptExecFilterNoHeader() OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--no-header"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// filter sample
func OptExecFilterSamples(nameSample ...string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"--samples", strings.Join(nameSample, ","), "--force-samples"}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// filter include
func OptExecFilterInclude(rule string) OptExec {
	return func(c *exec.Cmd) {
		args := []string{"-i", rule}
		c.Args = append(c.Args[:2], append(args, c.Args[2:]...)...)
	}
}

// Sort pathVcf by bcftools default arguments
func (s *Bcf) FilterCmd(argsFilter []string, opts ...OptExec) *exec.Cmd {
	cmdSlice := []string{"view", "--threads", strconv.Itoa(int(s.Threads))}
	cmdSlice = append(cmdSlice, argsFilter...)
	cmdFilter := exec.Command(s.PathBcf, cmdSlice...)
	return AppendOptExec(cmdFilter, opts...)
}

// clean samples
func (s *Bcf) CleanSamples(pathInput, pathOutput string) error {
	optsFilter := []OptExec{
		OptExecFilterSamples(),
		OptExecSetInputPath(pathInput),
		OptExecSetOutputPath(pathOutput),
	}
	cmdFilter := s.FilterCmd([]string{}, optsFilter...)
	return cmdFilter.Run()
}
