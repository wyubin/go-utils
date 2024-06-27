package bcf

import (
	"fmt"
	"os/exec"
)

type Tabix struct {
	PathExec string
}

func NewTabix() *Tabix {
	tabix := Tabix{
		PathExec: "tabix",
	}
	return &tabix
}

// index
func (s *Tabix) Index(idxSeq, idxBegin, idxEnd int, pathTrg string) error {
	cmd := exec.Command(s.PathExec, fmt.Sprintf("-s%d", idxSeq), fmt.Sprintf("-b%d", idxBegin), fmt.Sprintf("-e%d", idxEnd), pathTrg)
	return RunWithLog(cmd)
}
