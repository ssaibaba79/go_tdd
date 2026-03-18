package benchmark

import "strings"

func Repeat(c string, n int) (r string) {

	for i := 0; i < n; i++ {
		r = r + c
	}

	return
}

func RepeatOptimized(c string, n int) (r string) {
	var builder strings.Builder

	for i := 0; i < n; i++ {
		builder.WriteString(c)
	}
	return builder.String()
}
