package parser

import (
	"strconv"
	"strings"
)

type Instruction struct {
	iType string
	value int
}

func Parse(lines []string) (instructions []Instruction) {
	result := []Instruction{}
	for _, line := range lines {
		line := strings.Trim(line, "\n")
		if strings.HasPrefix(line, "@") {
			value, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			instruction := Instruction{"@", value}
			result = append(result, instruction)
		}
	}
	return result
}
