package utils

import "strings"

func SpacedStringToSlice(str string) (strSlice []string) {
	strs := strings.Split(str, " ")

	for _, rawStr := range strs {
		str := strings.TrimSpace(rawStr)
		strSlice = append(strSlice, str)
	}

	return strSlice
}
