package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	arr := []string{"A", "B", "C", "A", "C"}
	newArr := Uniq(arr)
	assert.Equal(t, 3, len(newArr))
}

func TestSubset(t *testing.T) {
	subset1 := []string{"A", "B", "C"}
	subset2 := []string{"A", "B", "C", "A", "C"}
	total := []string{"A", "B", "C", "D"}
	assert.True(t, Subset(subset1, total), "subset without repeat")
	assert.False(t, Subset(subset2, total), "subset with repeat")
}

func TestRemove(t *testing.T) {
	subset2 := []string{"A", "B", "C", "A", "C"}
	rmSet := Remove(subset2, "A")
	assert.Equal(t, []string{"B", "C", "C"}, rmSet, "remove repeat element")
}

func TestTranspose(t *testing.T) {
	arr := [][]string{
		{"A", "B", "C"},
		{"A", "B", "C"},
		{"A", "B", "C"},
	}
	res := Transpose(arr)
	assert.Equal(t, [][]string{
		{"A", "A", "A"},
		{"B", "B", "B"},
		{"C", "C", "C"},
	}, res)
}
