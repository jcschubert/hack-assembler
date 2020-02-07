package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/jcschubert/hack-assembler/hackparser"
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

// readLines takes a fileName and lines, opens a file, and writes
// the lines to it. The lines are separated by \n. If the file does
// not exist, it is created.
func writeLines(fileName string, lines []string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	for _, l := range lines {
		_, err := writer.WriteString(strings.Trim(l, "\n") + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func main() {

	if len(os.Args) != 2 {
		panic("Usage: hack-assembler <filename>")
	}
	fullPath := os.Args[1]

	path, name := getPathAndName(fullPath)

	lines, err := readLines(fullPath)
	if err != nil {
		panic(err)
	}

	instructions := hackparser.Parse(lines)
	binary := hackparser.Write(instructions)

	err = writeLines(path+name+".hack", binary)
	if err != nil {
		panic(err)
	}
}

func getPathAndName(path string) (folder string, name string) {
	parts := make([]string, 0)

	if strings.Contains(path, "/") {
		parts = strings.Split(path, "/")
		folder = strings.Join(parts[:len(parts)-1], "/") + "/"
		name = parts[len(parts)-1]
	} else {
		name = path
	}

	if strings.Contains(name, ".") {
		parts = strings.Split(name, ".")
		name = parts[0]
	} else {
		name = path
	}

	return folder, name
}
