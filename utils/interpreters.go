package utils

import "strings"

func SpacedStringToSlice(str string) (strSlice []string) {
	strSlice = strings.Split(str, " ")

	return strSlice
}
