package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetInputStdin() []string {
	in := bufio.NewReader(os.Stdin)
	var (
		input []string
	)

	for {
		var inpLine string
		_, err := fmt.Fscanln(in, &inpLine)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		input = append(input, inpLine)
	}
	return input
}
