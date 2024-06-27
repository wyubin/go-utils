package serialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSerialize(t *testing.T) {
	set := Set{}
	set.Add("a", "a", "b")
	assert.Equal(t, 2, len(set))
	err := set.UnmarshalText([]byte(`a,a,b,b,c,c,c`))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(set))
	assert.True(t, set.Contains("c"))
	set.Remove("c")
	assert.Equal(t, 2, len(set))
	assert.False(t, set.Contains("c"))
}
