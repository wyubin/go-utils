package bcf

import (
	"bytes"
	"context"
	"os/exec"

	"slices"

	"ailab.com/vcfgo/utils/vcf"
)

func (s *Bcf) GetHeader(pathSrc string, tagTypes ...string) ([]vcf.Header, error) {
	var res []vcf.Header
	bufHead := new(bytes.Buffer)
	cmdHead := exec.CommandContext(context.Background(), s.PathBcf, "head", pathSrc)
	cmdHead.Stdout = bufHead
	err := cmdHead.Run()
	if err != nil {
		return res, err
	}
	allTag, _ := vcf.ParseHeader(bufHead)
	for _, _info := range allTag {
		if len(tagTypes) == 0 || slices.Contains(tagTypes, _info.Tag) {
			res = append(res, _info)
		}
	}
	return res, err
}
