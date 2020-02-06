package parser

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
	i.dest = tokens[0]
	i.comp = tokens[1]
	i.jmp = ""
	return i
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

type Instruction interface {
	parse(code string) Instruction
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
		line := strings.Trim(line, "\n 	")

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
