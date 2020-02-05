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

func writeLines(fileName string, lines []string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	for _, l := range lines {
		i, err := writer.WriteString(l + "\n")
		fmt.Printf("%d bytes written", i)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func main() {
	lines, err := readLines("../asm/Add.asm")
	if err != nil {
		panic(err)
	}
	for i, l := range lines {
		fmt.Printf("%6d: %s", i, l)
	}

	err = writeLines("../asm/AddWritten.asm", lines)
	if err != nil {
		panic(err)
	}
}
