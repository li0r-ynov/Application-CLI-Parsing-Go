package main

import (
	"regexp"
)

func CorrigerPoncutation1(TexteCorriger string) string {
	regex := regexp.MustCompile(`\s+([.,!?;:])`)
	TexteCorriger = regex.ReplaceAllString(TexteCorriger, "$1")
	return TexteCorriger
}
