package main

import "strings"

func cap(mot string) string {
	if mot == "" {
		return ""
	}
	premierelettre := strings.ToUpper(mot[:1])
	restedumot := strings.ToLower(mot[1:])
	return premierelettre + restedumot
}
