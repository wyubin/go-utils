package maptool

import "fmt"

// combine args and check keys
func CheckCombineArgs(keysNeeded []string, args ...map[string]interface{}) (map[string]interface{}, error) {
	argCombine := map[string]interface{}{}
	Update(argCombine, args...)
	for _, key := range keysNeeded {
		if _, found := argCombine[key]; !found {
			return nil, fmt.Errorf("missing argument: %s", key)
		}
	}
	return argCombine, nil
}
