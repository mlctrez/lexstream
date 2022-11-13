package filter

import (
	"regexp"
)

func Regex(expression *regexp.Regexp, items []string) []string {
	var result []string
	for _, item := range items {
		if expression.MatchString(item) {
			result = append(result, item)
		}
	}
	return result
}
