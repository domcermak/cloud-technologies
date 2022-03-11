package common

import "sort"

func SortedKeys(attrs map[string]interface{}) []string {
	keys := []string{}
	for key := range attrs {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
