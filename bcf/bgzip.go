package bcf

import (
	"os/exec"
	"strconv"
)

type Bgzip struct {
	PathExec string
	Threads  int
}

type OptBgzip func(*Bgzip)

func SetBgzipThread(threads int) OptBgzip {
	return func(c *Bgzip) {
		c.Threads = threads
	}
}

func defaultBgzip() *Bgzip {
	return &Bgzip{
		PathExec: "bgzip",
		Threads:  2,
	}
}

func NewBgzip(opts ...OptBgzip) *Bgzip {
	o := defaultBgzip()
	for _, fn := range opts {
		fn(o)
	}
	return o
}

// compress
func (s *Bgzip) Compress(pathInput string) error {
	cmdIndex := exec.Command(s.PathExec, "-f", "--threads", strconv.Itoa(int(s.Threads)), pathInput)
	return RunWithLog(cmdIndex)
}

// pipe
func (s *Bgzip) CompressCmd(opts ...OptExec) *exec.Cmd {
	cmdCompress := exec.Command(s.PathExec, "-c", "--threads", strconv.Itoa(int(s.Threads)))
	return AppendOptExec(cmdCompress, opts...)
}
