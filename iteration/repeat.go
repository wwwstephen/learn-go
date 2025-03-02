package main

import (
	"strings"
)

func Repeat(v string, repeats int) string {
	var r strings.Builder
	for i := 0; i < repeats; i++ {
		r.WriteString(v)
	}
	return r.String()
}
