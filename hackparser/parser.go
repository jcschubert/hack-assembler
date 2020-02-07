package hackparser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type IInstruction struct {
	comp string
	dest string
	jmp  string
}

func (i IInstruction) parse(code string) Instruction {
	if strings.Contains(code, "=") {
		tokens := strings.Split(code, "=")
		i.dest = tokens[0]
		code = tokens[1]
	}

	if strings.Contains(code, ";") {
		tokens := strings.Split(code, ";")
		i.comp = tokens[0]
		i.jmp = tokens[1]
	} else {
		i.comp = code
	}

	return i
}

func (i IInstruction) write() string {
	comps := make(map[string]string)

	comps["0"] = "0101010"
	comps["1"] = "0111111"
	comps["-1"] = "0111010"
	comps["D"] = "0001100"
	comps["A"] = "0110000"
	comps["!D"] = "0001101"
	comps["!A"] = "0110001"
	comps["-D"] = "0001111"
	comps["-A"] = "0110011"
	comps["D+1"] = "0011111"
	comps["A+1"] = "0110111"
	comps["D-1"] = "0001110"
	comps["A-1"] = "0110010"
	comps["D+A"] = "0000010"
	comps["D-A"] = "0010011"
	comps["A-D"] = "0000111"
	comps["D&A"] = "0010101"
	comps["M"] = "1110000"
	comps["!M"] = "1110001"
	comps["-M"] = "1110011"
	comps["M+1"] = "1110111"
	comps["M-1"] = "1110010"
	comps["D+M"] = "1000010"
	comps["D-M"] = "1010011"
	comps["M-D"] = "1000111"
	comps["D&M"] = "1000000"
	comps["D|M"] = "1010101"

	dests := make(map[string]string)
	dests[""] = "000"
	dests["M"] = "001"
	dests["D"] = "010"
	dests["MD"] = "011"
	dests["A"] = "100"
	dests["AM"] = "101"
	dests["AD"] = "110"
	dests["AMD"] = "111"

	jumps := make(map[string]string)
	jumps[""] = "000"
	jumps["JGT"] = "001"
	jumps["JEQ"] = "010"
	jumps["JGE"] = "011"
	jumps["JLT"] = "100"
	jumps["JNE"] = "101"
	jumps["JLE"] = "110"
	jumps["JMP"] = "111"

	return "111" + comps[i.comp] + dests[i.dest] + jumps[i.jmp]
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
	return toBinary(i.value)
}

type Instruction interface {
	parse(code string) Instruction
	write() string
}

func stripComments(lines []string) (result []string) {
	for _, line := range lines {
		// Handle inline comments, then trim whitespace
		index := strings.IndexAny(line, "//")
		if index != -1 {
			line = line[:index]
		}

		result = append(result, line)
	}
	return result
}

func stripWhitespace(lines []string) (result []string) {
	for _, line := range lines {
		// Strip space, tab, newline
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	return result
}

var symbols = map[string]int{}

// only feed this with lines without comments or whitespace. only instructions!
func readSymbols(lines []string, symbols map[string]int) map[string]int {
	i := 0
	for _, line := range lines {
		// check if we have a label
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			// label of (LOOP) = LOOP
			label := line[1 : len(line)-1]
			// point to next instruction
			symbols[label] = i + 1
		} else {
			// inc instruction counter
			i++
		}
	}

	return symbols
}

func Parse(lines []string) []Instruction {
	instructions := []Instruction{}

	lines = stripComments(lines)
	lines = stripWhitespace(lines)
	readSymbols(lines, symbols)

	for _, line := range lines {
		if strings.HasPrefix(line, "@") {
			instruction := AInstruction{}
			instructions = append(instructions, instruction.parse(line))
		}
		if strings.ContainsAny(line, "=;") {
			instruction := IInstruction{}
			instructions = append(instructions, instruction.parse(line))
		}
	}

	return instructions
}

func Assemble(instruction Instruction) string {
	return instruction.write()
}

func Write(instructions []Instruction) []string {
	lines := []string{}

	for _, i := range instructions {
		lines = append(lines, i.write())
	}

	return lines
}

// toBinary converts an integer into its binary representation, stored as a string,
// padded with 0
func toBinary(value int) string {
	var digits []int
	for {
		if value == 1 {
			digits = append(digits, 1)
			break
		}
		if value == 0 {
			digits = append(digits, 0)
			break
		}

		remainder := value % 2

		if remainder == 1 {
			digits = append(digits, 1)
		} else {
			digits = append(digits, 0)
		}

		value = value / 2
	}

	var result bytes.Buffer

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 1 {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}
	return fmt.Sprintf("%016s", result.String())
}
