package utils

import (
	"bufio"
	"os"
)

func ScanString() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}

func ScanStringPerLine() (texts []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "q" {
			break
		}

		texts = append(texts, text)
	}

	return texts
}
