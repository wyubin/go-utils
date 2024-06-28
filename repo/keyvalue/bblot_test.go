package keyvalue

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/wyubin/go-utils/pathutils"
	"github.com/stretchr/testify/assert"
)

var (
	dirDB         string = filepath.Join(pathutils.DirExec(), "tmp")
	clientDefault *ClientBblot
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	os.MkdirAll(dirDB, os.ModePerm)
	clientDefault, _ = NewBblotClient(&OptionsBblot{
		Path:   filepath.Join(dirDB, "default.db"),
		Bucket: "0",
	})
	// set init key value
	clientDefault.Put([]byte("k1"), []byte("v1"))
}

func tearDown() {
	clientDefault.Close()
	os.RemoveAll(dirDB)
}

// test Uri2BblotOption
func TestUri2BblotOption(t *testing.T) {
	pathDB := filepath.Join(dirDB, "test.db")
	opt, err := Uri2BblotOption(fmt.Sprintf("bblot://%s?bucket=0&ro=true&timeout=1s", pathDB))
	assert.NoError(t, err)
	assert.Equal(t, pathDB, opt.Path, "Uri2BblotOption Pass")
	assert.Equal(t, "0", opt.Bucket, "Uri2BblotOption Pass")
	assert.Equal(t, true, opt.Options.ReadOnly, "Uri2BblotOption Pass")
	assert.Equal(t, 1*time.Second, opt.Options.Timeout, "Uri2BblotOption Pass")
}

func TestNewBblotClient(t *testing.T) {
	opt := &OptionsBblot{
		Path:   filepath.Join(dirDB, "testClient.db"),
		Bucket: "1",
	}
	client, err := NewBblotClient(opt)
	assert.NoError(t, err)
	assert.Equal(t, []byte("1"), client.bucketBytes, "NewBblotClient Pass")
}

func TestBblotPut(t *testing.T) {
	err := clientDefault.Put([]byte("k2"), []byte("v2"))
	assert.NoError(t, err)
	valTest, err := clientDefault.Get([]byte("k2"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("v2"), valTest, "BblotPut Pass")
}

func TestBblotGet(t *testing.T) {
	valTest, err := clientDefault.Get([]byte("k1"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("v1"), valTest, "BblotGet Pass")
}
