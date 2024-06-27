package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpl2Str(t *testing.T) {
	var strTmpl = `hello {{.name}}`
	var args = map[string]interface{}{
		"name": "world",
	}
	res, err := Tmpl2Str(strTmpl, args)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", res)
}
