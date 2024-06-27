package bcf

import (
	"path/filepath"
	"runtime"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	dirTest, _        = filepath.Abs(filepath.Join(filepath.Dir(filename), "test_file"))
)
