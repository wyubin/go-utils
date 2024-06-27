package vcf

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalVariant(t *testing.T) {
	strVariant := "chr1\t123\t.\tA\tG\t.\t.\ttagA=1;tagB=2\tGT:AD:DP:GQ:PL\t.\t./.:.:.:.:.\t1/1:0,2:2:.:49,6,0"
	variant, err := UnmarshalVariant(strVariant)
	assert.NoError(t, err)
	assert.Equal(t, "123", variant.POS)
	assert.Equal(t, "1", variant.INFO["tagA"])
	assert.Equal(t, "0,2", variant.Samples[2]["AD"])
	_, found := variant.Samples[2]["GQ"]
	assert.False(t, found)
}

func TestMarshalVariant(t *testing.T) {
	variant := NewVariant("chr1", "123", "A", "G")
	variant.INFO["tagA"] = "1"
	variant.Samples = append(variant.Samples, map[string]string{
		"GT": "1/1",
		"AD": "0,2",
	})
	strVariant := variant.Marshal()
	sliceVariant := strings.Split(strVariant, "\t")
	// fmt.Printf("strVariant:%+v\n", strVariant)
	assert.Equal(t, "chr1", strVariant[:4])
	assert.Equal(t, "tagA=1", sliceVariant[7])
}
