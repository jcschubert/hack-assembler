package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// readLines takes a fileName, splits its lines along '\n' and
// returns them
func readLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)
	lines := make([]string, 0)

	for {
		bytes, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		lines = append(lines, string(bytes))
	}

	return lines, nil
}

func main() {
	lines, err := readLines("main.go")
	if err != nil {
		panic(err)
	}
	for i, l := range lines {
		fmt.Printf("%6d: %s", i, l)
	}
}
