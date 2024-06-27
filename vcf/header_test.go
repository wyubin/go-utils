package vcf

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	pathTest, _       = filepath.Abs(filepath.Join(filepath.Dir(filename), "default_ex.vcf"))
)

func TestParseHeader(t *testing.T) {
	fileVcf, _ := os.Open(pathTest)
	defer fileVcf.Close()
	headers, _ := ParseHeader(fileVcf)
	assert.Equal(t, "fileformat", headers[0].ID)
	headerINFO := headers[5]
	assert.Equal(t, "##INFO=<ID=VEP_Symbol,Number=1,Type=String,Description=\"VEP SYMBOL\">", headerINFO.String())
}
