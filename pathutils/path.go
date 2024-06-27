package pathutils

import (
	"os"
	"path/filepath"
	"strings"
)

// get filename before "."
func Stem(pathFile string) string {
	filename := filepath.Base(pathFile)
	return filename[:strings.Index(filename, ".")]
}

// return self binary path
func DirExec() string {
	ex, _ := os.Executable()
	wd, _ := filepath.Abs(filepath.Dir(ex))
	return wd
}

// return all files in path
func GetAllFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// try to create dir
func MkDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
