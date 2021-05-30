package main

import (
	"strings"
)

func Contains(s []int, ID int) bool {
	for _, v := range s {
		if v == ID {
			return true
		}
	}

	return false
}

func Between(str string, before string, after string) string {
	a := strings.SplitAfterN(str, before, 2)
	b := strings.SplitAfterN(a[len(a)-1], after, 2)
	if 1 == len(b) {
		return b[0]
	}
	return b[0][0 : len(b[0])-len(after)]
}
