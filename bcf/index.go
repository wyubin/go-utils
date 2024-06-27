package bcf

import "os/exec"

// define Bcf operation processor
type Bcf struct {
	PathBcf  string
	Threads  int
	Compress bool
}

func defaultBcf() *Bcf {
	return &Bcf{
		PathBcf:  "bcftools",
		Threads:  2,
		Compress: true,
	}
}

// generate index of vcf.gz
func (s *Bcf) Index(pathVcf string) error {
	cmdIndex := exec.Command(s.PathBcf, "index", "-f", pathVcf)
	return RunWithLog(cmdIndex)
}

func NewBcf() *Bcf {
	return defaultBcf()
}
