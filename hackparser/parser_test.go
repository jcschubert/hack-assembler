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
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := Parse(c.input)
			if !reflect.DeepEqual(c.want, got) {
				t.Fatalf("Parse(%+v) should return %+v, but returned %+v", c.input, c.want, got)
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
		{0, "0"},
		{1, "1"},
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
