package app

import (
	"fmt"
	"strings"
)

func getValueByPath(data map[string]interface{}, path string) (interface{}, error) {
	keys := strings.Split(path, ".")
	var current interface{} = data
	for _, key := range keys {
		currMap, ok := current.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("key %s not found or parent is not a map", key)
		}
		val, exists := currMap[key]
		if !exists {
			return nil, fmt.Errorf("key %s not found", key)
		}
		current = val
	}
	return current, nil
}
