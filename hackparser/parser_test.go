package hackparser

import (
	"fmt"
	"reflect"
	"testing"
)

type ParserTestCase struct {
	desc  string
	input []string
	want  []Instruction
}

func TestParser(t *testing.T) {
	cases := []ParserTestCase{
		{
			"An empty program returns no instructions.",
			[]string{""},
			[]Instruction{},
		},
		{
			"A single @instruction with a value returns an @instruction with an address value.",
			[]string{"@1234"},
			[]Instruction{AInstruction{1234}},
		},
		{
			"A line comment does not result in an instruction.",
			[]string{"// This is a comment."},
			[]Instruction{},
		},
		{
			"An comment after an instruction is removed.",
			[]string{"@1234 // This is a comment."},
			[]Instruction{AInstruction{1234}},
		},
		{
			"An I-Instruction returns an I-Instruction",
			[]string{"D=A"},
			[]Instruction{
				IInstruction{dest: "D", comp: "A", jmp: ""},
			},
		},
		{
			"M=D+A is correctly parsed as an I-Instruction",
			[]string{"M=D+A"},
			[]Instruction{
				IInstruction{dest: "M", comp: "D+A", jmp: ""},
			},
		},
		{
			"M=D+A;JNE is correctly parsed as an I-Instruction",
			[]string{"M=D+A;JNE"},
			[]Instruction{
				IInstruction{dest: "M", comp: "D+A", jmp: "JNE"},
			},
		},
		{
			"0;JMP is correctly parsed as an I-Instruction",
			[]string{"0;JMP"},
			[]Instruction{
				IInstruction{dest: "", comp: "0", jmp: "JMP"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := Parse(c.input)
			if !reflect.DeepEqual(c.want, got) {
				t.Fatalf("Parse(%+q) should return %+q, but returned %+q", c.input, c.want, got)
			}
		})
	}
}

type ToBinaryTestCase struct {
	input int
	want  string
}

func TestToBinary(t *testing.T) {
	cases := []ToBinaryTestCase{
		{0, "0000000000000000"},
		{1, "0000000000000001"},
		{2, "0000000000000010"},
		{3, "0000000000000011"},
		{4, "0000000000000100"},
		{5, "0000000000000101"},
		{6, "0000000000000110"},
		{7, "0000000000000111"},
		{8, "0000000000001000"},
		{9, "0000000000001001"},
		{10, "0000000000001010"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d is converted to %v", c.input, c.want), func(t *testing.T) {
			got := toBinary(c.input)
			if got != c.want {
				t.Fatalf("ToBinary(%d) should return %q, but returned %q", c.input, c.want, got)
			}
		})
	}
}

type AssembleTestCase struct {
	instruction Instruction
	written     string
}

func TestAssemble(t *testing.T) {
	cases := []AssembleTestCase{
		{
			AInstruction{0},
			"0000000000000000",
		},
		{
			IInstruction{comp: "A", dest: "D", jmp: ""},
			"1110110000010000",
		},
	}

	// "111accccccdddjjj"

	for _, c := range cases {
		t.Run(fmt.Sprintf("%+v is assembled to %s", c.instruction, c.written), func(t *testing.T) {
			got := Assemble(c.instruction)
			if got != c.written {
				t.Fatalf("ToBinary(%+v) should return %q, but returned %q", c.instruction, c.written, got)
			}
		})
	}
}
