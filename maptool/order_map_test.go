package maptool

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestOrderMap(t *testing.T) {
	obj := `{"key3": {"TestVal":1}, "key2":  {"TestVal":2}, "key1":  {"TestVal":3}}`
	var o OrderedMap[itemTest]
	json.Unmarshal([]byte(obj), &o)
	assert.Equal(t, o.Order[0], "key3")
	byteO, _ := json.Marshal(o)
	assert.Equal(t, string(byteO)[2:6], "key3")
}

func TestOrderMapYaml(t *testing.T) {
	obj := `key3: {"TestVal":1}
key2:  {"TestVal":2}
key1:  {"TestVal":3}`
	var o OrderedMap[itemTest]
	err := yaml.Unmarshal([]byte(obj), &o)
	// fmt.Printf("o:%+v\n", o)
	assert.NoError(t, err)
	assert.Equal(t, o.Order[0], "key3")
}

type itemTest struct {
	Val int `json:"TestVal" yaml:"TestVal"`
}
