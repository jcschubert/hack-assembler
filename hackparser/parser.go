package hackparser

import (
	"strconv"
	"strings"
)

type IInstruction struct {
	comp string
	dest string
	jmp  string
}

func (i IInstruction) parse(code string) Instruction {
	tokens := strings.Split(code, "=")
	if len(tokens) == 2 {
		i.dest = tokens[0]
		i.comp = tokens[1]
	}
	i.jmp = ""
	return i
}

func (i IInstruction) write() string {
	return ""
}

type AInstruction struct {
	value int
}

func (i AInstruction) parse(code string) Instruction {
	value, err := strconv.Atoi(code[1:])
	if err != nil {
		panic(err)
	}
	i.value = value
	return i
}

func (i AInstruction) write() string {
	return ""
}

type Instruction interface {
	parse(code string) Instruction
	write() string
}

func Parse(lines []string) (instructions []Instruction) {
	result := []Instruction{}
	for _, line := range lines {
		// Handle inline comments, then trim whitespace
		index := strings.IndexAny(line, "//")
		if index != -1 {
			line = line[:index]
		}
		// Strip space, tab, newline
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handle @instructions
		if strings.HasPrefix(line, "@") {
			instruction := AInstruction{}
			result = append(result, instruction.parse(line))
		} else {
			instruction := IInstruction{}
			result = append(result, instruction.parse(line))
		}
	}

	return result
}

func Write(instructions []Instruction) []string {
	lines := []string{}

	for _, i := range instructions {
		lines = append(lines, i.write())
	}

	return lines
}