package maptool

import "golang.org/x/exp/constraints"

// return keys from map object
func Keys[T constraints.Ordered, P any](objMap map[T]P) []T {
	keys := []T{}
	for k := range objMap {
		keys = append(keys, k)
	}
	return keys
}

func Values[T constraints.Ordered, P any](objMap map[T]P) []P {
	values := []P{}
	for _, value := range objMap {
		values = append(values, value)
	}
	return values
}

// update keys value
func Update[T constraints.Ordered, P any](targetMap map[T]P, srcMaps ...map[T]P) {
	for _, srcMap := range srcMaps {
		for key, val := range srcMap {
			targetMap[key] = val
		}
	}
}

// pop key from map
func Pop[T constraints.Ordered, P any](targetMap map[T]P, key T) (P, bool) {
	v, ok := targetMap[key]
	if ok {
		delete(targetMap, key)
	}
	return v, ok
}

// copy map with single level
func Copy[T constraints.Ordered, P any](targetMap map[T]P) map[T]P {
	result := make(map[T]P)
	for k, v := range targetMap {
		result[k] = v
	}
	return result
}
