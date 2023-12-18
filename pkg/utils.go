package pkg

import (
	"strings"
)

func someHelperFunction(s string) []string {
	parts := strings.Split(s, ";")

	var result []string
	for _, part := range parts {
		subParts := strings.Split(part, ",")
		if len(result) > 0 {
			result[len(result)-1] += subParts[0]
			result = append(result, subParts[1:]...)
		} else {
			result = append(result, subParts...)
		}
	}
	return result
}
