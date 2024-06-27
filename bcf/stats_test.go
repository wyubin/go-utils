package bcf

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test for bcf Stats
func TestStats(t *testing.T) {
	bcf := NewBcf()
	res, err := bcf.Stats(filepath.Join(dirTest, "test.vcf"), ParserVariantCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, res["SNPCount"])
}
