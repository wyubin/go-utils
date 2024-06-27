package ver

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	dirTest, _        = filepath.Abs(filepath.Join(filepath.Dir(filename), "test_file"))
)

func TestVerParse(t *testing.T) {
	_, err := UseVersion([]string{"a.v0.0.1.txt", "b.v0.0.2.txt"}, "v0.0.3")
	assert.Error(t, err)
	verStr, err := UseVersionDir(dirTest, "v1.0.1")
	assert.NoError(t, err)
	assert.Equal(t, "abc.v1.0.1.tmpl", verStr)
	verLast, err := UseVersionDir(dirTest, "latest")
	assert.NoError(t, err)
	assert.Equal(t, "def.v1.1.0.xml", verLast)
}
