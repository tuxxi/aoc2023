package util

import (
	"bufio"
	"os"
)

func GetInputStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var (
		input []string
	)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
