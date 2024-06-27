package base

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValids(t *testing.T) {
	reg := regexp.MustCompile("^[0-9XYMT]+-[[:digit:]]+-[ATCG]+-[ATCG]+$")
	inputSlice := []string{"1-100-A-T", "MT-300-A-T", "X-300", "chrM-400-A-T"}
	res := Valids(reg, inputSlice...)
	assert.Equal(t, []string{"1-100-A-T", "MT-300-A-T"}, res)
}
