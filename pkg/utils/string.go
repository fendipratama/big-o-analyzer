// pkg/utils/string.go
package utils

import "strings"

func CountAny(input string, words ...string) int {
	count := 0
	for _, w := range words {
		count += strings.Count(input, w)
	}
	return count
}

func ContainsAny(input string, words ...string) bool {
	for _, w := range words {
		if strings.Contains(input, w) {
			return true
		}
	}
	return false
}
