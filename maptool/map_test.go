package maptool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	mapObj := map[string]struct{}{"A": {}, "B": {}, "D": {}}
	keys := Keys(mapObj)
	assert.Equal(t, 3, len(keys))
	assert.Contains(t, keys, "A")
	assert.NotContains(t, keys, "C")
}

func TestUpdate(t *testing.T) {
	mapObj := map[string]int{"A": 1, "B": 2, "D": 4}
	srcObj := map[string]int{"A": 0, "C": 3}
	Update(mapObj, srcObj)
	assert.Equal(t, 0, mapObj["A"])
	assert.Contains(t, mapObj, "C")
}

func TestPop(t *testing.T) {
	mapObj := map[string]int{"A": 1, "B": 2, "D": 4}
	val, _ := Pop(mapObj, "A")
	assert.Equal(t, 1, val)
	assert.NotContains(t, mapObj, "A")
}

func TestCopy(t *testing.T) {
	mapObj := map[string]int{"A": 1, "B": 2, "D": 4}
	copyObj := Copy(mapObj)
	copyObj["A"] = 0
	assert.Equal(t, 0, copyObj["A"])
	assert.Equal(t, 1, mapObj["A"])
}
