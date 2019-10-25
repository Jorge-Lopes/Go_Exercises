// Package word provide functions to deal with quotes.
package word

import "strings"

// UseCount return number of words in a quote.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return number of words in a quote.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
